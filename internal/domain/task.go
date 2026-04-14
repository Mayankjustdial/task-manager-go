package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	Pending    TaskStatus = "PENDING"
	InProgress TaskStatus = "IN_PROGRESS"
	Done       TaskStatus = "DONE"
)

var (
	ErrInvalidTitle = errors.New("title is required")
	ErrInvalidDate  = errors.New("due date must be in the future")
)

type Task struct {
	id          string
	title       string
	description string
	status      TaskStatus
	dueDate     time.Time
}

func NewTask(title, desc string, due time.Time) (*Task, error) {
	if title == "" {
		return nil, ErrInvalidTitle
	}
	if due.Before(time.Now()) {
		return nil, ErrInvalidDate
	}

	return &Task{
		id:          uuid.NewString(),
		title:       title,
		description: desc,
		status:      Pending,
		dueDate:     due,
	}, nil
}

func (t *Task) Update(title, desc *string, status *TaskStatus, due *time.Time) error {
	if title != nil && *title == "" {
		return ErrInvalidTitle
	}
	if due != nil && due.Before(time.Now()) {
		return ErrInvalidDate
	}

	if title != nil {
		t.title = *title
	}
	if desc != nil {
		t.description = *desc
	}
	if status != nil {
		t.status = *status
	}
	if due != nil {
		t.dueDate = *due
	}
	return nil
}

func (t *Task) ID() string          { return t.id }
func (t *Task) Title() string       { return t.title }
func (t *Task) Description() string { return t.description }
func (t *Task) Status() TaskStatus  { return t.status }
func (t *Task) DueDate() time.Time  { return t.dueDate }
