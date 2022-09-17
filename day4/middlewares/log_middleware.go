package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(e *echo.Echo) {
	log := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI: true, LogStatus: true,
		LogLatency: true, LogRequestID: true, LogMethod: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":     values.URI,
				"status":  values.Status,
				"latency": values.Latency,
				"method":  values.Method,
			}).Info("request")

			return nil
		},
	}))
}
