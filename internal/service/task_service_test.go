package service_test

import (
	"testing"
	"time"

	"task-manager/internal/domain"
	"task-manager/internal/repository"
	"task-manager/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	task, _ := svc.Create("title", "desc", time.Now().Add(time.Hour))

	result, err := svc.Get(task.ID())

	assert.NoError(t, err)
	assert.Equal(t, task.ID(), result.ID())
}

func TestGetTask_NotFound(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	_, err := svc.Get("invalid-id")

	assert.Error(t, err)
}

func TestUpdateTask(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	task, _ := svc.Create("title", "desc", time.Now().Add(time.Hour))

	newStatus := domain.InProgress

	updated, err := svc.Update(task.ID(), nil, nil, &newStatus, nil)

	assert.NoError(t, err)
	assert.Equal(t, domain.InProgress, updated.Status())
}

func TestUpdateTask_NotFound(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	newStatus := domain.Done

	_, err := svc.Update("invalid-id", nil, nil, &newStatus, nil)

	assert.Error(t, err)
}

func TestDeleteTask(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	task, _ := svc.Create("title", "desc", time.Now().Add(time.Hour))

	err := svc.Delete(task.ID())

	assert.NoError(t, err)
}

func TestDeleteTask_NotFound(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	err := svc.Delete("invalid-id")

	assert.Error(t, err)
}

func TestListTasks(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	svc.Create("task1", "desc", time.Now().Add(2*time.Hour))
	svc.Create("task2", "desc", time.Now().Add(1*time.Hour))

	tasks, err := svc.List(nil, 10, 0)

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)

	// Check sorting by due_date
	assert.True(t, tasks[0].DueDate().Before(tasks[1].DueDate()))
}

func TestListTasks_FilterByStatus(t *testing.T) {
	repo := repository.NewRepo()
	svc := service.New(repo)

	task1, _ := svc.Create("task1", "desc", time.Now().Add(time.Hour))

	status := domain.Done
	svc.Update(task1.ID(), nil, nil, &status, nil)

	filter := domain.Done
	tasks, err := svc.List(&filter, 10, 0)

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, task1.ID(), tasks[0].ID())
}
