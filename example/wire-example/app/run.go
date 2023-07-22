package app

import (
	"github.com/gin-gonic/gin"
	"kart-io/kart/example/wire-example/config"
	"kart-io/kart/example/wire-example/routers"
	"kart-io/kart/transports"
	kartHttp "kart-io/kart/transports/kart-http"
)

// Run 通过 Complete 函数完成
func Run(config *config.Config) error {
	// 实例化参数
	newServerConfig := kartHttp.NewServerConfig()
	// 重新给 Config 赋值
	if lastErr := config.InsecureServingOptions.ApplyTo(newServerConfig); lastErr != nil {
		return nil
	}

	httpServer := newServerConfig.Complete().New()

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

// RunV1 通过 With 方式处理
func RunV1(config *config.Config) error {
	// 实例化参数
	newServerConfig := kartHttp.NewServerConfig()
	// 重新给 Config 赋值
	if lastErr := config.InsecureServingOptions.ApplyTo(newServerConfig); lastErr != nil {
		return nil
	}

	handler := gin.Default()
	// 实例化 http 服务
	httpServer := kartHttp.NewServer(
		kartHttp.WithGinEngin(handler),
		kartHttp.WithInsecureServingInfo(newServerConfig.InsecureServing),
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
