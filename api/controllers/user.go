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

	isCreated, user := models.CreateUser(user)
	if isCreated {
		type ShortUser struct {
			Name   string `json:"name"`
			Mobile string `json:"mobile"`
			Role   string `json:"role"`
		}
		var shortUser ShortUser
		shortUser.Name = user.Name
		shortUser.Mobile = user.Mobile
		shortUser.Role = user.Role
		return c.JSON(fiber.Map{
			"success": isCreated,
			"user":    shortUser,
		})
	} else {
		return c.JSON(fiber.Map{
			"success": isCreated,
		})
	}
}
