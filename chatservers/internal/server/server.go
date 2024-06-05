package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	mw "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/middleware"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"time"
)

type Config struct {
	Port string
}

type Deps struct {
	Usecase usecases.UsecaseInterface
}

type Server struct {
	echo *echo.Echo
	cfg  Config
	Deps
}

func New(cfg Config, d Deps) (*Server, io.Closer) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(mw.Metrics)

	closer := jaegertracing.New(e, nil)

	e.Use(mw.Logging())
	e.Logger.SetLevel(log.DEBUG)

	server := &Server{
		echo: e,
		cfg:  cfg,
		Deps: d,
	}

	server.setRoutes()

	return server, closer
}

func (s *Server) setRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ChatServers Service")
	})

	s.echo.GET("/health", s.health)
	s.echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	s.echo.POST("/api/v1/chatservers", s.createChatServer)
	s.echo.GET("/api/v1/chatservers/:id", s.getChatServer)
	s.echo.GET("/api/v1/chatservers/search", s.searchChatServers)
	s.echo.GET("/api/v1/chatservers/of_user/:user_id", s.getUserChatServers)
	s.echo.POST("/api/v1/chatservers/:id/participants", s.createParticipant)
	s.echo.DELETE("/api/v1/chatservers/:id/participants/:participant_id", s.deleteParticipant)
	s.echo.POST("/api/v1/chatservers/:id/invites", s.createInvite)
}

func (s *Server) Start(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			s.echo.Logger.Printf("Server recovered from panic: %v", r)
		}
	}()

	go func() {
		if err := s.echo.Start(":" + s.cfg.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf(ctx, "server: %v", err)
		}
	}()

	// Wait until we receive a shutdown signal
	<-ctx.Done()

	logger.Info(ctx, "server: shutting down server gracefully")

	// Create a context with a 20-second timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Attempt a graceful shutdown
	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server: shutdown: %w", err)
	}

	logger.Info(ctx, "server: shutdown")

	return nil
}

func (s *Server) httpErrorMsg(err error) *models.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models.ErrorMessage{
		Message: err.Error(),
	}
}
