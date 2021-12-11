package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func sessionDetailsHandler(c *fiber.Ctx) error {
	loggedInUser := c.Locals("SessionUser")
	isLoggedIn := false
	if loggedInUser != nil {
		isLoggedIn = true
	}
	return c.JSON(&struct {
		LoggedIn bool
		User     interface{}
	}{
		LoggedIn: isLoggedIn,
		User:     loggedInUser,
	})
}
func sessionLoginHandler(c *fiber.Ctx) error {
	return nil
}
func sessionLogoutHandler(c *fiber.Ctx) error {
	return nil
}
