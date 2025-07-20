package user_handler

import (
	"go_grpc/clients/internal/config"
	grpc_conn "go_grpc/clients/internal/grpc"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type (
	IUserHandler interface {
		RegisterRoutes(group *echo.Group)
		CreateUser(c echo.Context) error
 	}

	UserHandlerDto struct {
		Conf   *config.Config
		Log    *slog.Logger
		GrpcServer *grpc_conn.GrpcClients
	}

	UserHandler struct {
		log *slog.Logger
		conf *config.Config
		grpcServer *grpc_conn.GrpcClients
	}
)

func  NewUserHandler(req UserHandlerDto) IUserHandler {
	return &UserHandler{log: req.Log, conf: req.Conf, grpcServer: req.GrpcServer}
}

func (handler *UserHandler) RegisterRoutes(group *echo.Group) {
	group.POST("/user", handler.CreateUser)
}





