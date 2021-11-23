package controller

import (
	"context"
	"go_ddd/internal/app/domain/model"
	"go_ddd/internal/app/usecase"
	"go_ddd/internal/pb"
)

type TaskController struct {
	pb.UnimplementedTaskServiceServer
	usecase usecase.TaskUsecase
}

func (c *TaskController) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task := c.usecase.Get(model.TaskId(in.Id))
	return convertPbTask(task), nil
}

func convertPbTask(task model.Task) *pb.GetTaskResponse {
	return &pb.GetTaskResponse{
		Id:      int64(task.Id),
		Title:   task.Title,
		Content: task.Content,
	}
}
