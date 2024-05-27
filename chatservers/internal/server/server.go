package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/usecases"
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
		return c.HTML(http.StatusOK, "ChatServers Service")
	})

	s.echo.GET("/health", s.health)
	s.echo.POST("/api/v1/chatservers", s.createChatServer)
	s.echo.GET("/api/v1/chatservers/:id", s.getChatServer)
	s.echo.GET("/api/v1/chatservers/search", s.searchChatServers)
	s.echo.GET("/api/v1/chatservers/of_user/:user_id", s.getUserChatServers)
	s.echo.POST("/api/v1/chatservers/:id/participants", s.createParticipant)
	s.echo.DELETE("/api/v1/chatservers/:id/participants/:participant_id", s.deleteParticipant)
	s.echo.POST("/api/v1/chatservers/:id/invites", s.createInvite)
}

func (s *Server) Start() {
	defer func() {
		if r := recover(); r != nil {
			s.echo.Logger.Printf("Server recovered from panic: %v", r)
		}
	}()

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
