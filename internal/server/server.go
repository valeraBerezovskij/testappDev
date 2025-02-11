package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"

	"example.com/taskservice/internal/delivery/rest"
	"example.com/taskservice/internal/repository"
	"example.com/taskservice/internal/service"
)

func NewServer(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	taskRepo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(taskRepo)

	handler := rest.NewHandler(taskService)

	r.POST("/tasks", handler.CreateTaskHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
