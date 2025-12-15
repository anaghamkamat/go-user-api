package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go-user-api/internal/logger"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		reqID, _ := c.Locals("requestId").(string)

		logger.Log.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("requestId", reqID),
			zap.Duration("duration", duration),
			zap.Int("status", c.Response().StatusCode()),
		)

		return err
	}
}
