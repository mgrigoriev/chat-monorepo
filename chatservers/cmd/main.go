package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/authclient"
	models2 "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var empty struct{}

func httpErrorMsg(err error) *models2.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models2.ErrorMessage{
		Message: err.Error(),
	}
}

func health(c echo.Context) error {
	response := struct{ Status string }{Status: "OK"}

	return c.JSON(http.StatusOK, response)
}

func authenticate(c echo.Context) (userID int, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("auth header not found")
	}

	authToken := strings.Split(authHeader, "Bearer ")
	if len(authToken) < 2 {
		return 0, fmt.Errorf("invalid token")
	}

	return authclient.Authenticate(authToken[1])
}

func createChatServer(c echo.Context) error {
	var request models2.CreateChatServerRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models2.CreateChatServerResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func getChatServer(c echo.Context) error {
	_, err := authenticate(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, httpErrorMsg(err))
	}

	id := c.Param("id")
	chatServerID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models2.GetChatServerResponse{
		ID:   int64(chatServerID),
		Name: "test",
	}
	return c.JSON(http.StatusOK, response)
}

func searchChatServers(c echo.Context) error {
	term := c.QueryParam("term")
	_ = term

	// ...

	response := []models2.GetChatServerResponse{
		{
			ID:   1,
			Name: "test",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func getUserChatServers(c echo.Context) error {
	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}
	_ = userID

	// ...

	response := []models2.GetChatServerResponse{
		{
			ID:   1,
			Name: "test",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func createParticipant(c echo.Context) error {
	var request models2.CreateParticipantRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models2.CreateParticipantResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func deleteParticipant(c echo.Context) error {
	participantIDParam := c.Param("participant_id")
	_, err := strconv.Atoi(participantIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func createInvite(c echo.Context) error {
	var request models2.CreateInviteRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models2.CreateInviteResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ChatServers Service")
	})

	e.GET("/health", health)

	e.POST("/api/v1/chatservers", createChatServer)
	e.GET("/api/v1/chatservers/:id", getChatServer)
	e.GET("/api/v1/chatservers/search", searchChatServers)
	e.GET("/api/v1/chatservers/of_user/:user_id", getUserChatServers)
	e.POST("/api/v1/chatservers/:id/participants", createParticipant)
	e.DELETE("/api/v1/chatservers/:id/participant/:participant_id", deleteParticipant)
	e.POST("/api/v1/chatservers/:id/invites", createInvite)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
