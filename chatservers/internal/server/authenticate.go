package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/logger"
	"github.com/mgrigoriev/goauth/authclient"
	"strings"
	"time"
)

const (
	authURL     = "http://users:8080/api/v1/users/auth"
	authTimeout = 5 * time.Second
)

func (s *Server) authenticate(c echo.Context) (user *authclient.CurrentUser, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("auth header not found")
	}

	authToken := strings.Split(authHeader, "Bearer ")
	if len(authToken) < 2 {
		return nil, fmt.Errorf("invalid token")
	}

	cfg := authclient.Config{AuthURL: authURL, Timeout: authTimeout}
	client := authclient.New(cfg, logger.Logger()) // TODO: Инициализировать при создании сервера, а не при каждом запросе

	return client.Authenticate(c.Request().Context(), authToken[1])
}
