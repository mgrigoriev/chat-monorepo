package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
)

func (s *Server) auth(c echo.Context) error {
	var request serverModels.AuthUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	user, err := s.Usecase.Auth(c.Request().Context(), models.AuthToken(*request.Token))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, s.httpErrorMsg(err))
	}

	response := serverModels.GetUserResponse{
		ID:             int64(user.ID),
		Email:          user.Email,
		Name:           user.Name,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}

	return c.JSON(http.StatusOK, response)
}