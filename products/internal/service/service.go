package service

import (
	"context"
	dataaccess "go_grpc/products/internal/data_access"
	pb "go_grpc/proto/products"
	"log/slog"
)

type ServiceDto struct {
	Log        *slog.Logger
	DataAccess *dataaccess.Queries
}

type Service struct {
	log        *slog.Logger
	dataAccess *dataaccess.Queries
}

type IService interface {
	GetAllProducts(ctx context.Context) (*pb.GetAllProductsResponse, error)
}

func NewService(req ServiceDto) IService {
	return &Service{log: req.Log, dataAccess: req.DataAccess}
}