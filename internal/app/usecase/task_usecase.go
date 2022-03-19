package usecase

import (
	"go_ddd/internal/app/domain/model"
	"go_ddd/internal/app/domain/repository"
)

type TaskUsecase interface {
	Get(model.TaskId) model.Task
}

type taskUsecase struct {
	repository repository.TaskRepository
}

func NewTaskUsecase(repository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repository: repository,
	}
}

func (us *taskUsecase) Get(id model.TaskId) model.Task {
	return us.repository.Get(id)
}
