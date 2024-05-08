package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/authclient"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const DefaultPort = "8080"

type Server struct {
	echo *echo.Echo
	port string
}

func NewServer() *Server {
	e := echo.New()

	s := &Server{
		echo: e,
		port: os.Getenv("PORT"),
	}

	s.configure()
	s.setRoutes()

	return s
}

func (s *Server) configure() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	s.echo.Logger.SetLevel(log.DEBUG)

	if s.port == "" {
		s.port = DefaultPort
	}
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
	s.echo.Logger.Fatal(s.echo.Start(":" + s.port))
}

func (s *Server) httpErrorMsg(err error) *models.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models.ErrorMessage{
		Message: err.Error(),
	}
}

func (s *Server) health(c echo.Context) error {
	response := struct{ Status string }{Status: "OK"}

	return c.JSON(http.StatusOK, response)
}

func (s *Server) authenticate(c echo.Context) (userID int, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("auth header not found")
	}

	authToken := strings.Split(authHeader, "Bearer ")
	if len(authToken) < 2 {
		return 0, fmt.Errorf("invalid token")
	}

	authClient := authclient.NewAuthClient()
	return authClient.Authenticate(authToken[1])
}

func (s *Server) createChatServer(c echo.Context) error {
	var request models.CreateChatServerRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.CreateChatServerResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func (s *Server) getChatServer(c echo.Context) error {
	userID, err := s.authenticate(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, s.httpErrorMsg(err))
	}

	// DEBUG
	c.Logger().Debug("Authenticated User ID: ", userID)

	id := c.Param("id")
	chatServerID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.GetChatServerResponse{
		ID:   int64(chatServerID),
		Name: "test",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) searchChatServers(c echo.Context) error {
	term := c.QueryParam("term")
	_ = term

	// ...

	response := []models.GetChatServerResponse{
		{
			ID:   1,
			Name: "test",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) getUserChatServers(c echo.Context) error {
	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}
	_ = userID

	// ...

	response := []models.GetChatServerResponse{
		{
			ID:   1,
			Name: "test",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) createParticipant(c echo.Context) error {
	var request models.CreateParticipantRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.CreateParticipantResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func (s *Server) deleteParticipant(c echo.Context) error {
	participantIDParam := c.Param("participant_id")
	_, err := strconv.Atoi(participantIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func (s *Server) createInvite(c echo.Context) error {
	var request models.CreateInviteRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.CreateInviteResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}
