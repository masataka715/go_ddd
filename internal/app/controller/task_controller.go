package controller

import (
	"context"
	"go_ddd/internal/pb"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer
}

func (s *TaskService) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	return &pb.GetTaskResponse{}, nil
}
