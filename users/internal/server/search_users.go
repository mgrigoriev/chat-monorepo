package server

import (
	"github.com/labstack/echo/v4"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
)

func (s *Server) searchUsers(c echo.Context) error {
	term := c.QueryParam("term")

	users, err := s.Usecase.SearchUsers(c.Request().Context(), term)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	var response []serverModels.GetUserResponse

	for _, user := range *users {
		response = append(response, serverModels.GetUserResponse{
			ID:             int64(user.ID),
			Email:          user.Email,
			Name:           user.Name,
			AvatarPhotoURL: user.AvatarPhotoURL,
		})
	}

	return c.JSON(http.StatusOK, response)
}
