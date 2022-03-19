package persistence

import (
	"go_ddd/internal/app/domain/model"
	"go_ddd/internal/app/domain/repository"

	"gorm.io/gorm"
)

type taskPersisitence struct {
	db *gorm.DB
}

func NewTaskPersistence(db *gorm.DB) repository.TaskRepository {
	return &taskPersisitence{
		db: db,
	}
}

func (*taskPersisitence) Get(id model.TaskId) model.Task {
	return model.Task{
		Id: id,
	}
}
