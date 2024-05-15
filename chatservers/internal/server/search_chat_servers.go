package server

import (
	"github.com/labstack/echo/v4"
	serverModels "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
)

func (s *Server) searchChatServers(c echo.Context) error {
	term := c.QueryParam("term")

	chatServers, err := s.Usecase.SearchChatServers(c.Request().Context(), term)
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
