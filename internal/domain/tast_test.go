package domain_test

import (
	"testing"
	"time"

	"task-manager/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewTask_Success(t *testing.T) {
	task, err := domain.NewTask("title", "desc", time.Now().Add(time.Hour))
	assert.NoError(t, err)
	assert.Equal(t, "title", task.Title())
}

func TestNewTask_Invalid(t *testing.T) {
	_, err := domain.NewTask("", "", time.Now().Add(time.Hour))
	assert.Error(t, err)
}
