package server

import (
	"encoding/json"
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
	"os"
	"strconv"
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
		return c.HTML(http.StatusOK, "Users Service")
	})

	s.echo.GET("/health", s.health)

	s.echo.POST("/api/v1/users", s.createUser)
	s.echo.GET("/api/v1/users/:id", s.getUser)
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

func (s *Server) createUser(c echo.Context) error {
	var request models.CreateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.CreateUserResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func (s *Server) getUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.GetUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	var request models.UpdateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.UpdateUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) login(c echo.Context) error {
	var request models.LoginUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.LoginUserResponse{
		Token: "example-token",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) auth(c echo.Context) error {
	var request models.AuthUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// TODO: Implement auth logic
	if *request.Token != "valid-token" {
		return c.JSON(http.StatusUnauthorized, s.httpErrorMsg(errors.New("invalid token")))
	}

	response := models.GetUserResponse{
		ID:    8,
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) searchUsers(c echo.Context) error {
	term := c.QueryParam("term")
	_ = term

	// ...

	response := []models.GetUserResponse{
		{
			ID:    1,
			Email: "test@mail.ru",
			Name:  "test",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) createFriendship(c echo.Context) error {
	var request models.CreateFriendshipRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := models.CreateFriendshipResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func (s *Server) getFriendshipList(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	response := []models.GetFriendshipResponse{
		{
			ID:         1,
			FollowerID: int64(userID),
			FollowedID: 2,
			Status:     "accepted",
		},
		{
			ID:         2,
			FollowerID: int64(userID),
			FollowedID: 3,
			Status:     "pending",
		},
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Server) acceptFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func (s *Server) declineFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func (s *Server) deleteFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}
