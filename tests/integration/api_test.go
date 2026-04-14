package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"task-manager/internal/handler"
	"task-manager/internal/repository"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	repo := repository.NewRepo()
	svc := service.New(repo)
	h := handler.New(svc)
	h.Register(r)

	return r
}

func createTaskHelper(r *gin.Engine) (string, *httptest.ResponseRecorder) {
	body := map[string]interface{}{
		"title":    "test task",
		"due_date": time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	return resp["id"].(string), w
}

func TestCreateTask(t *testing.T) {
	r := setupRouter()

	_, w := createTaskHelper(r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateTask_Invalid(t *testing.T) {
	r := setupRouter()

	body := `{"title": "", "due_date": "2026-04-20T10:00:00Z"}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTask(t *testing.T) {
	r := setupRouter()

	id, _ := createTaskHelper(r)

	req, _ := http.NewRequest("GET", "/tasks/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTask_NotFound(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks/invalid-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateTask(t *testing.T) {
	r := setupRouter()

	id, _ := createTaskHelper(r)

	body := `{"status":"IN_PROGRESS"}`
	req, _ := http.NewRequest("PUT", "/tasks/"+id, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask_NotFound(t *testing.T) {
	r := setupRouter()

	body := `{"status":"DONE"}`
	req, _ := http.NewRequest("PUT", "/tasks/invalid-id", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteTask(t *testing.T) {
	r := setupRouter()

	id, _ := createTaskHelper(r)

	req, _ := http.NewRequest("DELETE", "/tasks/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteTask_NotFound(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/invalid-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListTasks(t *testing.T) {
	r := setupRouter()

	createTaskHelper(r)
	createTaskHelper(r)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListTasks_WithFilter(t *testing.T) {
	r := setupRouter()

	createTaskHelper(r)

	req, _ := http.NewRequest("GET", "/tasks?status=PENDING", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
