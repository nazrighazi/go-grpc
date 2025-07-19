package route

import (
	"go_grpc/clients/internal/config"
	"go_grpc/clients/internal/handlers"
	"log/slog"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Route struct {
	Router *echo.Echo
	Conf   *config.Config
	Log    *slog.Logger
	GrpcServer *grpc.ClientConn
}

func APIRoute(route *Route) {

	group := route.Router.Group("api/v1")

	handlers := handlers.NewHandlers(handlers.HandlersDto{Router: route.Router, Conf: route.Conf, Log: route.Log, GrpcServer: route.GrpcServer})

	// Declare routes for handlers
	handlers.ProductsHandler.RegisterRoutes(group)
}
