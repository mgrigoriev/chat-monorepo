package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func JaegerTracing(tracer opentracing.Tracer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

			var span opentracing.Span
			if err != nil {
				span = tracer.StartSpan(req.URL.Path)
			} else {
				span = tracer.StartSpan(req.URL.Path, opentracing.ChildOf(spCtx))
			}
			defer span.Finish()

			// Setting additional tags
			ext.HTTPMethod.Set(span, req.Method)
			ext.HTTPUrl.Set(span, req.URL.String())

			// Set the context with span
			ctx := opentracing.ContextWithSpan(c.Request().Context(), span)
			c.SetRequest(c.Request().WithContext(ctx))

			c.Set("Tracer", tracer)
			c.Set("Span", span)

			if err := next(c); err != nil {
				c.Error(err)
				ext.Error.Set(span, true)
				span.LogKV("error", err.Error())
			}

			return nil
		}
	}
}
