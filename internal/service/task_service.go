package service

import (
	"task-manager/internal/domain"
	"task-manager/internal/repository"
	"time"
)

type TaskService struct{ repo repository.TaskRepository }

func New(r repository.TaskRepository) *TaskService {
	return &TaskService{r}
}

func (s *TaskService) Create(title, desc string, due time.Time) (*domain.Task, error) {
	t, err := domain.NewTask(title, desc, due)
	if err != nil {
		return nil, err
	}
	return t, s.repo.Create(t)
}

func (s *TaskService) Get(id string) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(id string, title, desc *string, status *domain.TaskStatus, due *time.Time) (*domain.Task, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if err := t.Update(title, desc, status, due); err != nil {
		return nil, err
	}
	return t, s.repo.Update(t)
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *TaskService) List(status *domain.TaskStatus, limit, offset int) ([]*domain.Task, error) {
	return s.repo.List(status, limit, offset)
}
