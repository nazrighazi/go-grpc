package service

import (
	"context"
	"errors"
	"fmt"
	pb "go_grpc/proto/users"

	dataaccess "go_grpc/users/internal/data_access"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) RegisterUser(ctx context.Context, req *pb.RegisterUserRequestDto) (*pb.RegisterUserResponseDto, error) {

	fmt.Println("here")
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		fmt.Println("here 1")
		return nil, errors.New("failed to hash password")
	}
    
	_ , err = s.dataAccess.CreateUser(ctx, dataaccess.CreateUserParams{
		Email: req.Email,
		PasswordHash: string(hashedBytes),
		Role: pgtype.Text{
			String: "user",
			Valid: true,
		},
	})
	if err != nil { 
		fmt.Println("here 2")
		s.log.Error("error creating user by email", "error", err)
		return nil, errors.New("error creating user")
	}

	return &pb.RegisterUserResponseDto{
		Status: "success",
		Message: "user created successfully",
	},nil
}
