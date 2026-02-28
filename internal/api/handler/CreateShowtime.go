package handler

import (
	"cinema-ticket/internal/model"
	"cinema-ticket/internal/service"
	"errors"

	"github.com/gofiber/fiber/v2"
)


func CreateShowTime(c *fiber.Ctx) error {
	var input model.CreateShowTimeRequest
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	showtime, err := service.CreateShowTime(input)
	if errors.Is(err, service.ErrSessionNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "session not found, please create the session",
		})
	}

	if errors.Is(err, service.ErrMovieNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "movie not found, please create the movie",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "showtime has been created successfully",
		"showTime-id": showtime.ID,
	})
}