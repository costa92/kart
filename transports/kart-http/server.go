package kart_http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"kart-io/kart/transports/kart-http/middlewares"

	"github.com/costa92/errors"
	"github.com/costa92/logger"
)

type Server struct {
	Config          *HttpConfig
	GinEngin        *gin.Engine
	httpServer      *http.Server
	middlewares     []string
	healthz         bool
	enableMetrics   bool
	enableProfiling bool
}

func NewServer(opts ...Option) *Server {
	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}
	initAPIServer(srv)
	return srv
}

func initAPIServer(s *Server) {
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

func (s *Server) Setup() {
	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.Infow("gin endpoint setup ", "httpMethod", httpMethod, "absolutePath",
			absolutePath, "handlerName", handlerName, "nuHandlers", nuHandlers)
	}
}

func (s *Server) InstallMiddlewares() {
	if s.Config == nil {
		return
	}
	for _, m := range s.Config.Middlewares {
		mw, ok := middlewares.Middlewares[m]
		if !ok {
			continue
		}
		s.GinEngin.Use(mw)
	}
}

func (s *Server) InstallAPIs() {
	// Healthz 检测健康
	if s.healthz {
		s.GinEngin.GET("/healthz", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, http.StatusText(http.StatusOK))
		})
	}
	// install metric handler
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.GinEngin)
	}

	// install pprof handler
	if s.enableProfiling {
		pprof.Register(s.GinEngin)
	}
}

func (s *Server) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorw("appService recover err", "err", err)
		}
	}()
	serverConfig := s.Config
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConfig.Port),
		Handler:        s.GinEngin,
		ReadTimeout:    time.Duration(serverConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Infow("start run http server", "port", serverConfig.Port)
	if err := s.httpServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) { // 如果是关闭状态，不当异常处理
			log.Print("start run failed server:", serverConfig.Port)
			return err
		}
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	logger.Infow("[HTTP] server stopping")
	return s.httpServer.Shutdown(ctx)
}
