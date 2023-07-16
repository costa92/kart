package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"kart-io/kart/transports"
	kart_grpc "kart-io/kart/transports/kart-grpc"
	kartHttp "kart-io/kart/transports/kart-http"
)

var ProviderHttpSeverSet = wire.NewSet(NewConfig, NewHttpSever)

type App struct {
	GenericAPIServer *transports.GenericAPIServer
}

func NewConfig() *kartHttp.HttpConfig {
	return &kartHttp.HttpConfig{
		Port:          "8080",
		Healthz:       true,
		EnableMetrics: true,
		Name:          "kart",
	}
}

func NewHttpSever(config *kartHttp.HttpConfig, engine *gin.Engine) (*App, error) {
	gs, err := initSever(config, engine)
	if err != nil {
		return nil, err
	}
	return &App{
		GenericAPIServer: gs,
	}, nil
}

func initSever(config *kartHttp.HttpConfig, handler *gin.Engine) (*transports.GenericAPIServer, error) {
	// 实例化 http 服务
	httpServer := kartHttp.NewServer(
		kartHttp.WithGinEngin(handler),
		kartHttp.WithConfig(config),
	)

	grcConfig := &kart_grpc.GrpcConfig{
		Port: "8081",
		Addr: "0.0.0.0",
	}
	grpcServer := kart_grpc.NewGrpcServer(kart_grpc.WithConfig(grcConfig))
	// 运行 http 与 rpc
	gs := transports.NewGenericAPIServer(
		transports.Server(
			httpServer,
			grpcServer,
		),
		transports.Name(config.Name),
	)
	return gs, nil
}

func (a *App) Run() error {
	return a.GenericAPIServer.Run()
}
