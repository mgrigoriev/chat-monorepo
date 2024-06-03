package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) getChatServer(c echo.Context) error {
	if user, err := s.authenticate(c); err != nil {
		return c.JSON(http.StatusUnauthorized, s.httpErrorMsg(err))
	} else {
		c.Logger().Debug("Authenticated User ID: ", user.ID)
	}

	id := c.Param("id")
	chatServerID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	chatServer, err := s.Usecase.GetChatServerByID(c.Request().Context(), models.ChatServerID(chatServerID))
	if err != nil {
		return c.JSON(http.StatusNotFound, s.httpErrorMsg(err))
	}

	response := serverModels.GetChatServerResponse{
		ID:   int64(chatServer.ID),
		Name: chatServer.Name,
	}

	return c.JSON(http.StatusOK, response)
}
