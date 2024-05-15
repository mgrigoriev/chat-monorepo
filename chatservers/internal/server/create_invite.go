package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) createInvite(c echo.Context) error {
	var request serverModels.CreateInviteRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	chatServerID, err := strconv.Atoi(c.Param("chatserver_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	userID := models.UserID(*request.UserID)

	inviteID, err := s.Usecase.CreateInvite(c.Request().Context(), models.ChatServerID(chatServerID), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.CreateInviteResponse{ID: int64(inviteID)}
	return c.JSON(http.StatusCreated, response)
}
