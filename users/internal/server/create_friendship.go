package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) createFriendship(c echo.Context) error {
	var request serverModels.CreateFriendshipRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	followerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	friendshipID, err := s.Usecase.CreateFriendship(c.Request().Context(), models.UserID(followerID), models.UserID(*request.UserID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.CreateFriendshipResponse{ID: int64(friendshipID)}
	return c.JSON(http.StatusCreated, response)
}
