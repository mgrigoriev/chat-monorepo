package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	authclient "github.com/mgrigoriev/chat-monorepo/chatservers/internal/services/auth"
	"strings"
)

func (s *Server) authenticate(c echo.Context) (userID int, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("auth header not found")
	}

	authToken := strings.Split(authHeader, "Bearer ")
	if len(authToken) < 2 {
		return 0, fmt.Errorf("invalid token")
	}

	authClient := authclient.New()
	return authClient.Authenticate(authToken[1])
}
