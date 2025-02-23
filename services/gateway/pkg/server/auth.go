package server

import (
	"gateway/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type TestStruct struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func (s *Server) sign_up(c *fiber.Ctx) error {
	var input models.User
	err := c.BodyParser(&input)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if !input.IsValid() {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, "invalid data")
	}
	err = s.services.UserService.CreateUser(input)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	tokens, err := s.services.UserService.SignInByEP(input.Email, input.Password)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(tokens)
}

func (s *Server) sign_in(c *fiber.Ctx) error {
	var input models.User
	err := c.BodyParser(&input)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if !input.IsValid() {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, "invalid data")
	}
	err = s.services.UserService.CreateUser(input)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	tokens, err := s.services.UserService.SignInByEP(input.Email, input.Password)
	if err != nil {
		return s.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(tokens)
}
