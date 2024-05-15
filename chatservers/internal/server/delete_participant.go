package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/internal/models"
	"net/http"
	"strconv"
)

func (s *Server) deleteParticipant(c echo.Context) error {
	participantIDParam := c.Param("participant_id")
	participantID, err := strconv.Atoi(participantIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	err = s.Usecase.DeleteParticipant(c.Request().Context(), models.ParticipantID(participantID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	return c.NoContent(http.StatusOK)
}
