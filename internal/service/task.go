package service

import (
	"time"

	"example.com/taskservice/internal/domain"
	"example.com/taskservice/internal/repository"
)

type TaskService interface {
	CreateTask(title, description string) (*domain.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(title, description string) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}

	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}
