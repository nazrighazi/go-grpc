package service

import (
	"context"
	pb "go_grpc/proto/users"
	dataaccess "go_grpc/users/internal/data_access"
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
	LoginUser(ctx context.Context, req *pb.LoginUserRequestDto) (*pb.LoginUserResponseDto, error)
	RegisterUser(ctx context.Context, req *pb.RegisterUserRequestDto) (*pb.RegisterUserResponseDto, error)
}

func NewService(req ServiceDto) IService {
	return &Service{log: req.Log, dataAccess: req.DataAccess}
}