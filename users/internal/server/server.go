package server

import (
	"fmt"
	"go_grpc/users/internal/config"
	"go_grpc/users/internal/route"
	"log"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

type HttpServerDto struct {
	Conf *config.Config
	Log  *slog.Logger
	Db *pgx.Conn
}

type httpServer struct {
	conf  *config.Config
	log   *slog.Logger
	db    *pgx.Conn
}

func NewServer(req HttpServerDto) *httpServer {
	return &httpServer{conf: req.Conf, log: req.Log, db: req.Db}
}


func (httpServer *httpServer) Start() {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", httpServer.conf.Server.Host, httpServer.conf.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	route.APIRoute(&route.Route{
		Conf:   httpServer.conf,
		Log:    httpServer.log,
		GrpcServer: grpcServer,
		Listener: listener,
		Db: httpServer.db,
	})
}

