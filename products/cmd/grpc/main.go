package main

import (
	"context"
	"fmt"
	"go_grpc/products/internal/config"
	"go_grpc/products/internal/database"
	"go_grpc/products/internal/logger"
	"go_grpc/products/internal/server"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := logger.LoadLogger(config)
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(context.Background(), config)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("Error pinging database:", err)
    	panic(err)
	}

	defer db.Close(context.Background())

	server.NewServer(server.HttpServerDto{Conf: config, Log: logger, Db: db}).Start()

	
}