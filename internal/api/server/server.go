package server

import (
	"cinema-ticket/config"
	"cinema-ticket/internal/api/handler"

	"github.com/gofiber/fiber/v2"
)

func Start(cfg config.ServerConfig) {

	app := fiber.New()

	app.Post("/session", handler.CreateSession) 
	app.Post("/session/showtime", handler.CreateShowTime)
	app.Post("/session/movie", handler.CreateMovie)
	app.Post("/session/room", handler.CreateRoom)

	app.Listen(cfg.Address())

}