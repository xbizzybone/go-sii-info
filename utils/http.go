package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNextMiddleWare(c *fiber.Ctx) error {
	request_id := uuid.New().String()
	ctx := context.WithValue(c.Context(), "request_id", request_id)
	c.SetUserContext(ctx)

	return c.Next()
}
