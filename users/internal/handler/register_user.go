package handler

import (
	"context"
	pb "go_grpc/proto/users"
)

func (handler *UserHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequestDto) (*pb.RegisterUserResponseDto, error) {

	response, err := handler.service.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}