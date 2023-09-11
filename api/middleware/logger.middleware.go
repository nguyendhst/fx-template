package middleware

import (
	"fmt"
	"time"

	v4 "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nguyendhst/lagile/module/logger"
)

func LoggerMiddleware(lg logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				var traceID string
				if traceID = c.Request().Header.Get("X-Tracer-Id"); traceID != "" {
					c.Set("trace_id", traceID)
				} else {
					traceID = fmt.Sprintf("%d-%s", time.Now().Unix(), v4.New().String())
					c.Set("trace_id", traceID)
				}

				lg.LogRequest(
					"trace_id", traceID,
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
				)

				return nil
			},
		},
	)
}
