package repository

import (
	"errors"
	"sort"
	"sync"
	"task-manager/internal/domain"
)

type InMemoryRepo struct {
	store map[string]*domain.Task
	mu    sync.RWMutex
}

func NewRepo() *InMemoryRepo {
	return &InMemoryRepo{store: map[string]*domain.Task{}}
}

func (r *InMemoryRepo) Create(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[t.ID()] = t
	return nil
}

func (r *InMemoryRepo) GetByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.store[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return t, nil
}

func (r *InMemoryRepo) Update(t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[t.ID()] = t
	return nil
}

func (r *InMemoryRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.store[id]; !ok {
		return errors.New("not found")
	}
	delete(r.store, id)
	return nil
}

func (r *InMemoryRepo) List(status *domain.TaskStatus, limit, offset int) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []*domain.Task
	for _, t := range r.store {
		if status != nil && t.Status() != *status {
			continue
		}
		tasks = append(tasks, t)
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].DueDate().Before(tasks[j].DueDate()) })
	if offset > len(tasks) {
		return []*domain.Task{}, nil
	}
	end := offset + limit
	if end > len(tasks) {
		end = len(tasks)
	}
	return tasks[offset:end], nil
}
