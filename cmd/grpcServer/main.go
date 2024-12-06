package main

import (
	"database/sql"
	"net"

	"github.com/Aberos/fullcycle-grpc/internal/database"
	"github.com/Aberos/fullcycle-grpc/internal/pb"
	"github.com/Aberos/fullcycle-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	gprcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(gprcServer, categoryService)
	reflection.Register(gprcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := gprcServer.Serve(lis); err != nil {
		panic(err)
	}
}
