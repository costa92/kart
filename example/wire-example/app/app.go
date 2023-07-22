package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"kart-io/kart/example/wire-example/config"
	"kart-io/kart/example/wire-example/options"
	"kart-io/kart/example/wire-example/routers"
	"kart-io/kart/internal/command"
	"kart-io/kart/transports"
	kartHttp "kart-io/kart/transports/kart-http"
)

const commandDesc = `The Kart API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/costa92/kart`

var ProviderHttpSeverSet = wire.NewSet(NewConfig)

func NewConfig() *command.App {
	opts := options.NewOptions()
	a := command.NewApp(
		"kart",
		// command.WithNoConfig(),
		command.WithOptions(opts),
		command.WithDescription(commandDesc),
		command.WithRunFunc(run(opts)),
	)
	return a
}

func run(opts *options.Options) command.RunFunc {
	return func(basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}

func Run(config *config.Config) error {
	serverConfig := &kartHttp.HttpConfig{
		Healthz: config.ServerRunOption.Healthz,
		Name:    config.ServerRunOption.Name,
		Port:    config.InsecureServingOptions.BindPort,
	}
	handler := gin.Default()
	// 实例化 http 服务
	httpServer := kartHttp.NewServer(
		kartHttp.WithGinEngin(handler),
		kartHttp.WithConfig(serverConfig),
		kartHttp.WithMiddlewares(config.ServerRunOption.Middlewares),
		kartHttp.WithEnableMetrics(config.FeatureOptions.EnableMetrics),
		kartHttp.WithEnableProfiling(config.FeatureOptions.EnableProfiling),
	)

	// 添加路由
	routers.InitRoute(httpServer.GinEngin)

	// 实例化服务
	gs := transports.NewGenericAPIServer(
		transports.Server(
			httpServer,
		),
	)

	// 运行服务
	return gs.Run()
}
