package handler

import (
	"go_ddd/internal/app/controller"
	"go_ddd/internal/app/infrastructure/persistence"
	"go_ddd/internal/app/usecase"
	"go_ddd/internal/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterTask(s *grpc.Server, db *gorm.DB) {
	persistence := persistence.NewTaskPersistence(db)
	usecase := usecase.NewTaskUsecase(persistence)
	controller := controller.NewTaskController(usecase)
	pb.RegisterTaskServiceServer(s, controller)
}
