package handler

import (
	"strconv"

	"task-manager/internal/domain"
	"task-manager/internal/dto"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.TaskService
}

func New(s *service.TaskService) *Handler {
	return &Handler{s}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/tasks", h.Create)
	r.GET("/tasks/:id", h.Get)
	r.PUT("/tasks/:id", h.Update)
	r.DELETE("/tasks/:id", h.Delete)
	r.GET("/tasks", h.List)
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	t, err := h.svc.Create(req.Title, req.Description, req.DueDate)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(201, dto.ToResponse(t))
}

func (h *Handler) Get(c *gin.Context) {
	t, err := h.svc.Get(c.Param("id"))
	if err != nil {
		c.JSON(404, "not found")
		return
	}
	c.JSON(200, dto.ToResponse(t))
}

func (h *Handler) Update(c *gin.Context) {
	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	t, err := h.svc.Update(c.Param("id"), req.Title, req.Description, req.Status, req.DueDate)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, dto.ToResponse(t))
}

func (h *Handler) Delete(c *gin.Context) {
	if err := h.svc.Delete(c.Param("id")); err != nil {
		c.JSON(404, "not found")
		return
	}
	c.Status(204)
}

func (h *Handler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	var status *domain.TaskStatus
	if s := c.Query("status"); s != "" {
		t := domain.TaskStatus(s)
		status = &t
	}
	tasks, _ := h.svc.List(status, limit, offset)
	var res []dto.TaskResponse
	for _, t := range tasks {
		res = append(res, dto.ToResponse(t))
	}
	c.JSON(200, res)
}
