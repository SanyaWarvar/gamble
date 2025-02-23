package server

import "github.com/gofiber/fiber/v2"

func (s *Server) NewErrorResponse(c *fiber.Ctx, status_code int, error_message string) error {
	//TODO add log
	c.Status(status_code)
	return c.JSON(fiber.Map{"details": error_message})
}
