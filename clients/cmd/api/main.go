package main

import (
	"go_grpc/clients/internal/config"
	"go_grpc/clients/internal/logger"
	"go_grpc/clients/internal/server"
)

func main() {

	// Load config
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Load logger
	logger, err := logger.LoadLogger(config)
	if err != nil {
		panic(err)
	}

	// Load server
	server.NewServer(server.HttpServerDto{Conf: config, Log: logger}).Start()





}