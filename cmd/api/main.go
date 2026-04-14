package main

import (
	"task-manager/internal/handler"
	"task-manager/internal/repository"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := repository.NewRepo()
	svc := service.New(repo)
	h := handler.New(svc)
	h.Register(r)
	r.Run(":8080")
}
