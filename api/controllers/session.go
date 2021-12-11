package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func sessionDetailsHandler(c *fiber.Ctx) error {
	loggedInUser := c.Locals("SessionUser")
	return c.JSON(loggedInUser)
}
func sessionLoginHandler(c *fiber.Ctx) error {
	return nil
}
func sessionLogoutHandler(c *fiber.Ctx) error {
	return nil
}
