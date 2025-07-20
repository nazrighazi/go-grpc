package service

import (
	"context"
	"errors"
	pb "go_grpc/proto/users"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) LoginUser(ctx context.Context, req *pb.LoginUserRequestDto) (*pb.LoginUserResponseDto, error) {

	// 1. Get user by email
	user, err := s.dataAccess.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Error("error getting user by email", "error", err)
		return nil, errors.New("user not found")
	}

	// 2. Verify password hash
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil,errors.New("invalid credentials")
	}


	return nil,nil
}
