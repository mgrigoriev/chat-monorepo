package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"github.com/mgrigoriev/chat-monorepo/users/internal/usecases"
	"net/http"
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

func New(ctx context.Context, cfg Config, d Deps) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

func (s *Server) Start() {
	s.echo.Logger.Fatal(s.echo.Start(":" + s.cfg.Port))
}

func (s *Server) httpErrorMsg(err error) *models.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models.ErrorMessage{
		Message: err.Error(),
	}
}
