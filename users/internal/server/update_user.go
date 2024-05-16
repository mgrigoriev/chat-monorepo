package server

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	models "github.com/mgrigoriev/chat-monorepo/users/internal/models"
	serverModels "github.com/mgrigoriev/chat-monorepo/users/internal/server/models"
	"net/http"
	"strconv"
)

func (s *Server) updateUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	var request serverModels.UpdateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, s.httpErrorMsg(err))
	}

	user := models.User{
		Name:           request.Name,
		Email:          request.Email,
		Password:       request.Password,
		AvatarPhotoURL: request.AvatarPhotoURL,
	}

	updatedUser, err := s.Usecase.UpdateUser(c.Request().Context(), models.UserID(userID), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, s.httpErrorMsg(err))
	}

	response := serverModels.UpdateUserResponse{
		ID:             int64(updatedUser.ID),
		Email:          updatedUser.Email,
		Name:           updatedUser.Name,
		AvatarPhotoURL: updatedUser.AvatarPhotoURL,
	}

	return c.JSON(http.StatusOK, response)
}
