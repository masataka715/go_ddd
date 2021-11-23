package persistence

import (
	"go_ddd/internal/app/domain/model"
	"go_ddd/internal/app/domain/repository"
)

type taskPersisitence struct {
}

func NewTaskPersistence() repository.TaskRepository {
	return &taskPersisitence{}
}

func (*taskPersisitence) Get(id model.TaskId) model.Task {
	return model.Task{
		Id: id,
	}
}
