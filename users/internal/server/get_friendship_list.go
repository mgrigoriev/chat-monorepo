package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) getFriendshipList(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	friendships, err := s.Usecase.GetFriendshipList(c.Request().Context(), models.UserID(userID))
	if err != nil {
		return err
	}

	var response []serverModels.GetFriendshipResponse

	for _, friendship := range *friendships {
		response = append(response, serverModels.GetFriendshipResponse{
			ID:         int64(friendship.ID),
			FollowerID: int64(friendship.FollowerID),
			FollowedID: int64(friendship.FollowedID),
			Status:     friendship.Status,
		})
	}

	return c.JSON(http.StatusOK, response)
}
