package httpserver

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nguyendhst/lagile/api/middleware"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/module/logger"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Configs *config.Config
	Logger  logger.Logger
}

type Server struct {
	Server       *echo.Echo
	ServerConfig *ServerConfig `name:"server"`
	Prefix       *echo.Group
}

type ServerConfig struct {
	ServerAddress string `name:"server_address"`
	RateLimit     int    `name:"rate_limit"`
	Logger        logger.Logger
}

func (s *Server) Start() error {
	s.ServerConfig.Logger.Info("Starting Echo server")

	data, err := json.MarshalIndent(s.Server.Routes(), "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile("routes.json", data, 0o644)

	err = s.Server.Start(s.ServerConfig.ServerAddress)

	s.ServerConfig.Logger.Fatal("Error starting Echo server")
	return err
}

func (s *Server) ApplyMiddleware(middleware ...echo.MiddlewareFunc) {
	s.Server.Use(middleware...)
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.Server.Shutdown(ctx)
	if err != nil {
		s.ServerConfig.Logger.Error("Failed to gracefully shutdown echo server")
	} else {
		s.ServerConfig.Logger.Info("Echo server gracefully stopped")
	}

	return err
}

func (s *Server) SetPrefix(prefix string) *Server {
	group := s.Server.Group(prefix)
	return &Server{Server: s.Server, ServerConfig: s.ServerConfig, Prefix: group}
}

func (s *Server) Group(prefix string) *echo.Group {
	return s.Prefix.Group(prefix)
}

func (s *Server) GET(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	s.Prefix.GET(path, handler, middleware...)
}

func NewEchoServer(p Params) *Server {
	server := echo.New()

	// Add middleware
	server.Use(
		middleware.InvalidPathResponseFormatMiddleware,
		middleware.LoggerMiddleware(p.Logger),
	)

	rateLimit := 0
	if p.Configs.Env.App.Server.RateLimit.Enabled {
		rateLimit = p.Configs.Env.App.Server.RateLimit.Max
	}

	return &Server{Server: server, ServerConfig: &ServerConfig{
		ServerAddress: p.Configs.Env.App.Server.Address,
		RateLimit:     rateLimit,
		Logger:        p.Logger,
	}}
}
