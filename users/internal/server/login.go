package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
)

func (s *Server) login(c echo.Context) error {
	var request models.LoginUserRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	token, err := s.Usecase.Login(c.Request().Context(), *request.Email, *request.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, s.httpErrorMsg(err))
	}

	response := models.LoginUserResponse{
		Token: string(token),
	}
	return c.JSON(http.StatusOK, response)
}
