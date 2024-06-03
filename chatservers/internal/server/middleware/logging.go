package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mgrigoriev/chat-monorepo/chatservers/pkg/logger"
)

func Logging() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logFields := []interface{}{
				"URI", v.URI,
				"status", v.Status,
				"method", v.Method,
				"remote_ip", c.Request().RemoteAddr,
			}

			ctx := c.Request().Context()

			if v.Status >= 500 {
				logger.ErrorKV(ctx, "Request", logFields...)
			} else {
				logger.InfoKV(ctx, "Request", logFields...)
			}

			return nil
		},
	})
}
