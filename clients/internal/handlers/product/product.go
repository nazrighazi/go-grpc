package product_handler

import (
	"go_grpc/clients/internal/config"
	grpc_conn "go_grpc/clients/internal/grpc"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type (
	IProductHandler interface {
		RegisterRoutes(group *echo.Group)
		FindAll(c echo.Context) error
	}

	ProductHandlerDto struct {
		Conf   *config.Config
		Log    *slog.Logger
		GrpcServer *grpc_conn.GrpcClients
	}

	ProductHandler struct {
		log *slog.Logger
		conf *config.Config
		grpcServer *grpc_conn.GrpcClients
	}
)

func  NewProductHandler(req ProductHandlerDto) IProductHandler {
	return &ProductHandler{log: req.Log, conf: req.Conf, grpcServer: req.GrpcServer}
}

func (handler *ProductHandler) RegisterRoutes(group *echo.Group) {
	group.GET("/products", handler.FindAll)
}





