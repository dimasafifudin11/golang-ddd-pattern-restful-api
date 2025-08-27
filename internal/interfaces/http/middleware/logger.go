package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		latency := stop.Sub(start)

		log.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency_ms": latency.Milliseconds(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}).Info("Request processed")

		return err
	}
}
