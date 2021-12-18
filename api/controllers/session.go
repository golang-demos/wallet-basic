package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
	"go.mongodb.org/mongo-driver/bson"
)

func sessionDetailsHandler(c *fiber.Ctx) error {
	// loggedInUser := c.Locals("SessionUser")

	sessToken := string(c.Request().Header.Peek("SESS-TOKEN"))
	var user models.User
	if sessToken != "" {
		database.UserCollection.FindOne(context.Background(), bson.M{"token": sessToken}).Decode(&user)
	}
	isLoggedIn := false
	if !user.ID.IsZero() {
		isLoggedIn = true
	}
	return c.JSON(&struct {
		LoggedIn bool
		User     models.User
	}{
		LoggedIn: isLoggedIn,
		User:     user.ToShort(),
	})
}

func sessionLoginHandler(c *fiber.Ctx) error {
	var postData models.UserLoginData
	c.BodyParser(&postData)

	errors := postData.Validate()
	if errors != nil {
		return c.JSON(errors)
	}

	isCorrect, token := models.Login(postData.Mobile, postData.Password)
	if isCorrect {
		c.Set("SESS-TOKEN", token)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"isLoggedIn": isCorrect,
			"token":      token,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"isLoggedIn": isCorrect,
			"token":      token,
		})
	}
}
func sessionLogoutHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
