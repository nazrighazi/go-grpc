package grpc_conn

import (
	"fmt"

	"go_grpc/clients/internal/config"
	products "go_grpc/proto/products"
	users "go_grpc/proto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Product products.ProductsClient
	User    users.UsersClient
}

func NewGrpcClients(conf *config.Config) (*GrpcClients, error) {
	productConn, err := grpc.NewClient(fmt.Sprintf("%s:%d", conf.ProductService.Host, conf.ProductService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to products service: %w", err)
	}

	userConn, err := grpc.NewClient(fmt.Sprintf("%s:%d", conf.UserService.Host, conf.UserService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to users service: %w", err)
	}

	productClient := products.NewProductsClient(productConn)
	userClient := users.NewUsersClient(userConn)

	return &GrpcClients{
		Product: productClient,
		User:    userClient,
	}, nil
}
