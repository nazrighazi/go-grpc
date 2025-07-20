package handler

import (
	"context"
	"errors"
	pb "go_grpc/proto/users"
)

func (handler *UserHandler) LoginUser(ctx context.Context, req *pb.LoginUserRequestDto) (*pb.LoginUserResponseDto, error) {

	response, err := handler.service.LoginUser(ctx, req)
	if err != nil {
		return nil, errors.New("failed to get all products")
	}

	return response, nil
}