package handlers

import (
	"go_grpc/clients/internal/config"
	grpc_conn "go_grpc/clients/internal/grpc"
	product_handler "go_grpc/clients/internal/handlers/product"
	user_handler "go_grpc/clients/internal/handlers/user"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type (

	HandlersDto struct {
		Router *echo.Echo
		Conf   *config.Config
		Log    *slog.Logger
		GrpcServer *grpc_conn.GrpcClients
	}

	Handlers struct {
		ProductHandler product_handler.IProductHandler
		UserHandler user_handler.IUserHandler
	}

)

func NewHandlers(req HandlersDto) *Handlers {
	return &Handlers{
		ProductHandler: product_handler.NewProductHandler(product_handler.ProductHandlerDto{Conf: req.Conf, Log: req.Log, GrpcServer: req.GrpcServer}),
		UserHandler: user_handler.NewUserHandler(user_handler.UserHandlerDto{Conf: req.Conf, Log: req.Log, GrpcServer: req.GrpcServer}),

	}
}