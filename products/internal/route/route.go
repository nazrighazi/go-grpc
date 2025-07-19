package route

import (
	"fmt"
	"go_grpc/products/internal/config"
	dataaccess "go_grpc/products/internal/data_access"
	"go_grpc/products/internal/handler"
	"go_grpc/products/internal/service"
	pb "go_grpc/proto/products"
	"log"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

type Route struct {
	Conf *config.Config
	Log  *slog.Logger
	GrpcServer *grpc.Server
	Listener net.Listener
	Db *pgx.Conn
}

func APIRoute(route *Route) {
	
	dataaccess := dataaccess.NewDataAccess(dataaccess.DataAccessDto{Db: route.Db, Log: route.Log})

	service := service.NewService(service.ServiceDto{Log: route.Log, DataAccess: dataaccess})

	handler := handler.NewProductHandler(handler.ProductHandlerDto{Service: service})

	pb.RegisterProductsServer(route.GrpcServer, handler)

	fmt.Print("Product service is running on port :50051...")

	if err := route.GrpcServer.Serve(route.Listener); err != nil {
		log.Fatal("failed to serve: &v", err)
	}
}