package repository

import "task-manager/internal/domain"

type TaskRepository interface {
	Create(*domain.Task) error
	GetByID(string) (*domain.Task, error)
	Update(*domain.Task) error
	Delete(string) error
	List(*domain.TaskStatus, int, int) ([]*domain.Task, error)
}
