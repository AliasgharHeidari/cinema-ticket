package server

import (
	"cinema-ticket/internal/api/handler"

	"github.com/gofiber/fiber/v2"
)

func Start() {

	app := fiber.New()

	app.Post("/session", handler.CreateSession) 





}