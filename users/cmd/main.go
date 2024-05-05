package main

import (
	"encoding/json"
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mgrigoriev/chat-monorepo/users/server/models"
	"net/http"
	"os"
	"strconv"
)

var empty struct{}

func httpErrorMsg(err error) *models.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models.ErrorMessage{
		Message: err.Error(),
	}
}

func createUser(c echo.Context) error {
	var request models.CreateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.CreateUserResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.GetUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	var request models.UpdateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.UpdateUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func login(c echo.Context) error {
	var request models.LoginUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.LoginUserResponse{
		Token: "example-token",
	}
	return c.JSON(http.StatusOK, response)
}

func auth(c echo.Context) error {
	var request models.AuthUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// TODO: Implement auth logic
	if *request.Token != "valid-token" {
		return c.JSON(http.StatusUnauthorized, httpErrorMsg(errors.New("invalid token")))
	}

	response := models.GetUserResponse{
		ID:    8,
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func searchUsers(c echo.Context) error {
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

func createFriendship(c echo.Context) error {
	var request models.CreateFriendshipRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.CreateFriendshipResponse{ID: 1}
	return c.JSON(http.StatusCreated, response)
}

func getFriendshipList(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
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

func acceptFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func declineFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func deleteFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	_, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	return c.NoContent(http.StatusOK)
}

func health(c echo.Context) error {
	response := struct{ Status string }{Status: "OK"}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Users Service")
	})

	e.GET("/health", health)

	e.POST("/api/v1/users", createUser)
	e.GET("/api/v1/users/:id", getUser)
	e.PUT("/api/v1/users/:id", updateUser)

	e.POST("/api/v1/users/login", login)
	e.POST("/api/v1/users/auth", auth)

	e.GET("/api/v1/users/search", searchUsers)

	e.POST("/api/v1/users/:id/friendships", createFriendship)
	e.GET("/api/v1/users/:id/friendships", getFriendshipList)
	e.PUT("/api/v1/users/:id/friendships/:friendship_id/accept", acceptFriendship)
	e.PUT("/api/v1/users/:id/friendships/:friendship_id/decline", declineFriendship)
	e.DELETE("/api/v1/users/:id/friendships/:friendship_id", deleteFriendship)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
