package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/chatservers/internal/server/models"
	"net/http"
)

func (s *Server) createChatServer(c echo.Context) error {
	var request serverModels.CreateChatServerRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	chatserver := models.ChatServer{
		UserID: 1, // TODO: Define after auth is implemented
		Name:   *request.Name,
	}

	chatserverID, err := s.Usecase.CreateChatServer(c.Request().Context(), chatserver)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.CreateChatServerResponse{ID: int64(chatserverID)}
	return c.JSON(http.StatusCreated, response)
}
