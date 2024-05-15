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

func (s *Server) createParticipant(c echo.Context) error {
	var request serverModels.CreateParticipantRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	chatServerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	participantID, err := s.Usecase.CreateParticipant(c.Request().Context(), models.ChatServerID(chatServerID), models.UserID(*request.UserID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.CreateParticipantResponse{ID: int64(participantID)}
	return c.JSON(http.StatusCreated, response)
}
