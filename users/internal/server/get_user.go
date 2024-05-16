package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) getUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	user, err := s.Usecase.GetUserByID(c.Request().Context(), models.UserID(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, s.httpErrorMsg(err))
	}

	response := serverModels.GetUserResponse{
		ID:             int64(user.ID),
		Email:          user.Email,
		Name:           user.Name,
		AvatarPhotoURL: user.AvatarPhotoURL,
	}

	return c.JSON(http.StatusOK, response)
}
