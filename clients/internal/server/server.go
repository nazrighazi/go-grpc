package server

import (
	"fmt"
	"go_grpc/clients/internal/config"
	"go_grpc/clients/internal/route"
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HttpServerDto struct {
	Conf  *config.Config
	Log   *slog.Logger
}
type httpServer struct {
	conf  *config.Config
	log   *slog.Logger
	router *echo.Echo
}

func NewServer(req HttpServerDto) *httpServer {
	echoApp := echo.New()
	return &httpServer{conf: req.Conf, log: req.Log, router: echoApp}
}

func (httpServer *httpServer) Start() {

	httpServer.router.Use(middleware.Recover())
	httpServer.router.Use(middleware.Logger())

	// Api for health check
	httpServer.router.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// Connect to product service
	productConn, err := grpc.NewClient(":50051",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()


	httpServer.InitializeHandlers(productConn)

	serverUrl := fmt.Sprintf("%s:%d", httpServer.conf.Server.Host, httpServer.conf.Server.Port)
	httpServer.router.Logger.Fatal(httpServer.router.Start(serverUrl))
}

func (httpServer *httpServer) InitializeHandlers(grpc *grpc.ClientConn) {
	route.APIRoute(&route.Route{
		Router: httpServer.router,
		Conf:   httpServer.conf,
		Log:    httpServer.log,
		GrpcServer: grpc,
	})
}
