package handlers

import (
	"go_grpc/clients/internal/config"
	products_handler "go_grpc/clients/internal/handlers/products"
	"log/slog"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type (

	HandlersDto struct {
		Router *echo.Echo
		Conf   *config.Config
		Log    *slog.Logger
		GrpcServer *grpc.ClientConn
	}

	Handlers struct {
		ProductsHandler products_handler.IProductsHandler
	}

)

func NewHandlers(req HandlersDto) *Handlers {
	return &Handlers{
		ProductsHandler: products_handler.NewProductsHandler(products_handler.ProductsHandlerDto{Conf: req.Conf, Log: req.Log, GrpcServer: req.GrpcServer}),
	}
}