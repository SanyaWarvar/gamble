package server

import (
	"gateway/pkg/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	services services.Services
}

func NewServer(s services.Services) *Server {
	return &Server{services: s}
}

func (s *Server) CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}:${port}] ${time} ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05 02-Jan-2006",
		TimeZone:   "Asia/Krasnoyarsk",
	}))

	auth := app.Group("/auth")
	{
		auth.Post("/sign-up", s.sign_up)
		auth.Post("/sign-in", s.sign_in)
	}

	return app
}

func (s *Server) Run(port string) {
	app := s.CreateApp()

	log.Fatal(app.Listen(":" + port))
}
