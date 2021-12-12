package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-demos/ecommerce-basic/models"
)

func userSignupHandler(c *fiber.Ctx) error {

	var postData models.UserSignupData
	c.BodyParser(&postData)

	user := new(models.User).Init(postData)

	errors := user.Validate()
	if errors != nil {
		return c.JSON(errors)
	}

	result := models.CreateUser(user)
	return c.JSON(result)
}
