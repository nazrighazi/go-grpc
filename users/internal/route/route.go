package route

import (
	"fmt"
	pb "go_grpc/proto/users"
	"go_grpc/users/internal/config"
	dataaccess "go_grpc/users/internal/data_access"
	"go_grpc/users/internal/handler"
	"go_grpc/users/internal/service"
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

	handler := handler.NewUserHandler(handler.UserHandlerDto{Service: service})

	pb.RegisterUsersServer(route.GrpcServer, handler)

	fmt.Print("Product service is running on port :50052...")

	if err := route.GrpcServer.Serve(route.Listener); err != nil {
		log.Fatal("failed to serve: &v", err)
	}
}