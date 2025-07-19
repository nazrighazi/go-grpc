package products_handler

import (
	"go_grpc/clients/internal/config"
	pb "go_grpc/proto/products"
	"log/slog"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type (
	IProductsHandler interface {
		RegisterRoutes(group *echo.Group)
		FindAll(c echo.Context) error
	}

	ProductsHandlerDto struct {
		Conf   *config.Config
		Log    *slog.Logger
		GrpcServer *grpc.ClientConn
	}

	ProductsHandler struct {
		log *slog.Logger
		conf *config.Config
		grpcServer pb.ProductsClient
	}
)

func  NewProductsHandler(req ProductsHandlerDto) IProductsHandler {
	return &ProductsHandler{log: req.Log, conf: req.Conf, grpcServer: pb.NewProductsClient(req.GrpcServer)}
}

func (handler *ProductsHandler) RegisterRoutes(group *echo.Group) {
	group.GET("/products", handler.FindAll)
}





