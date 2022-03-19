package main

import (
	"go_ddd/internal/app/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	dsn := "root:root@tcp(127.0.0.1:3306)/go_ddd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	handler.RegisterTask(s, db)

	log.Printf("gRPC server listening on " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}
