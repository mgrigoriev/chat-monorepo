package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"time"
)

const (
	appName   = "chatservers"
	namespace = appName
)

var ms struct {
	responseTimeHistogram *prometheus.HistogramVec
}

func init() {
	ms.responseTimeHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "http",
			Name:      appName + "_histogram_response_time_seconds",
			Help:      "Время ответа от сервера",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
		},
		[]string{"method", "status_code", "is_error"},
	)
}

func Metrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Выполняем следующий обработчик в цепочке
		err := next(c)

		// Измеряем время выполнения
		duration := time.Since(start)
		statusCode := c.Response().Status

		// Определяем, была ли ошибка
		isError := err != nil

		// Собираем метрики
		responseTimeHistogramObserve(c.Request().Method, statusCode, isError, duration)

		return err
	}
}

func responseTimeHistogramObserve(method string, statusCode int, isError bool, d time.Duration) {
	ms.responseTimeHistogram.WithLabelValues(method, strconv.Itoa(statusCode), strconv.FormatBool(isError)).Observe(d.Seconds())
}
