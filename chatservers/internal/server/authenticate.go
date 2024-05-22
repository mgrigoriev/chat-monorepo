package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mgrigoriev/goauth/authclient"
	"net/http"
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

	httpClient := &http.Client{Timeout: authTimeout}
	cfg := authclient.Config{AuthURL: authURL}
	client := authclient.New(cfg, httpClient) // TODO: Инициализировать при создании сервера, а не при каждом запросе

	return client.Authenticate(authToken[1])
}
