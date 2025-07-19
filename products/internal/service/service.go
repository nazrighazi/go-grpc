package service

import (
	"context"
	dataaccess "go_grpc/products/internal/data_access"
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
	GetAllProducts(ctx context.Context) ([]dataaccess.Product, error)
}

func NewService(req ServiceDto) IService {
	return &Service{log: req.Log, dataAccess: req.DataAccess}
}