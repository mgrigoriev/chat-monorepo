package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	"net/http"
	"strconv"
)

func (s *Server) acceptFriendship(c echo.Context) error {
	friendshipIDParam := c.Param("friendship_id")
	friendshipID, err := strconv.Atoi(friendshipIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	err = s.Usecase.AcceptFriendship(c.Request().Context(), models.FriendshipID(friendshipID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	return c.NoContent(http.StatusOK)
}
