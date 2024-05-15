package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) health(c echo.Context) error {
	response := struct{ Status string }{Status: "OK"}

	return c.JSON(http.StatusOK, response)
}
