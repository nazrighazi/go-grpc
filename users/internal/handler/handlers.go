package handler

import (
	pb "go_grpc/proto/users"
	service "go_grpc/users/internal/service"
)
type UserHandlerDto struct {
	Service service.IService
}
type UserHandler struct {
	pb.UnimplementedUsersServer
	service service.IService
}

func NewUserHandler(req UserHandlerDto) *UserHandler {
	return &UserHandler{service: req.Service}
}