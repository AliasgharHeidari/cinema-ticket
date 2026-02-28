package handler

import (
	"cinema-ticket/internal/model"
	"cinema-ticket/internal/service"
	"errors"

	"github.com/gofiber/fiber/v2"
)



func CreateMovie(c *fiber.Ctx) error {
	var input model.CreateMovieRequest
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	movie, err := service.CreateMovie(input)
	if errors.Is(err, service.ErrInternal) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal error, please try again later",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "movie created successfully",
		"movie-id": movie.ID,
	})
}