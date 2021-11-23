package repository

import "go_ddd/internal/app/domain/model"

type TaskRepository interface {
	Get(model.TaskId) model.Task
}
