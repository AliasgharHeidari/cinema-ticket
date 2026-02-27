package handler

import (
	"cinema-ticket/internal/model"

	"github.com/gofiber/fiber/v2"
)

func CreateSession(c *fiber.Ctx) error {
	var Input model.Session
	err := c.BodyParser(Input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid request body",
		})
	}

	resp, err := service.CreateSession


}
