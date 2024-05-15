package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) getUserChatServers(c echo.Context) error {
	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	chatServers, err := s.Usecase.GetUserChatServers(c.Request().Context(), models.UserID(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, s.httpErrorMsg(err))
	}

	var response []serverModels.GetChatServerResponse

	for _, chatServer := range *chatServers {
		response = append(response, serverModels.GetChatServerResponse{
			ID:   int64(chatServer.ID),
			Name: chatServer.Name,
		})
	}

	return c.JSON(http.StatusOK, response)
}
