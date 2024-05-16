package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
)

func (s *Server) createUser(c echo.Context) error {
	var request serverModels.CreateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	user := models.User{
		Name:           *request.Name,
		Email:          *request.Email,
		Password:       *request.Password,
		AvatarPhotoURL: request.AvatarPhotoURL,
	}

	userID, err := s.Usecase.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.CreateUserResponse{ID: int64(userID)}
	return c.JSON(http.StatusCreated, response)
}
