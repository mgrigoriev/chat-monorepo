package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/users/pkg/logger"
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

func New(cfg Config, d Deps) *Server {
	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logFields := []interface{}{
				"URI", v.URI,
				"status", v.Status,
				"method", v.Method,
				"remote_ip", c.Request().RemoteAddr,
			}

			if v.Status >= 500 {
				logger.ErrorKV(c.Request().Context(), "Server Error", logFields...)
			} else {
				logger.InfoKV(c.Request().Context(), "Request", logFields...)
			}

			return nil
		},
	}))

	e.Logger.SetLevel(log.DEBUG)

	s := &Server{
		echo: e,
		cfg:  cfg,
		Deps: d,
	}

	s.setRoutes()

	return s
}

func (s *Server) setRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Users Service")
	})

	limiter1 := middleware.NewRateLimiterMemoryStore(1)   // 1 request / sec
	limiter10 := middleware.NewRateLimiterMemoryStore(10) // 10 requests / sec

	// Test rate limiter for the method (bash):
	//   for i in `seq 1 100`; do curl --location 'localhost:8080/health'; done
	s.echo.GET("/health", s.health, middleware.RateLimiter(limiter10))

	// Test rate limiter for the method (bash):
	// 	 for i in `seq 1 100`; do curl --location 'localhost:8080/api/v1/users/1'; done
	s.echo.GET("/api/v1/users/:id", s.getUser, middleware.RateLimiter(limiter1))

	s.echo.POST("/api/v1/users", s.createUser)
	s.echo.PUT("/api/v1/users/:id", s.updateUser)
	s.echo.POST("/api/v1/users/login", s.login)
	s.echo.POST("/api/v1/users/auth", s.auth)
	s.echo.GET("/api/v1/users/search", s.searchUsers)
	s.echo.POST("/api/v1/users/:id/friendships", s.createFriendship)
	s.echo.GET("/api/v1/users/:id/friendships", s.getFriendshipList)
	s.echo.PUT("/api/v1/users/:id/friendships/:friendship_id/accept", s.acceptFriendship)
	s.echo.PUT("/api/v1/users/:id/friendships/:friendship_id/decline", s.declineFriendship)
	s.echo.DELETE("/api/v1/users/:id/friendships/:friendship_id", s.deleteFriendship)
}

func (s *Server) Start(ctx context.Context) error {
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
