package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
)

func apiV1Handler(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	sessToken := string(c.Request().Header.Peek("SESS-TOKEN"))
	var user models.User
	if sessToken != "" {
		database.UserCollection.FindOne(context.Background(), bson.M{"token": sessToken}).Decode(&user)
	}
	c.Locals("SESSION_USER_ID", user.ID)
	return c.Next()
}

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", apiV1Handler)

	// session
	// Check User session
	v1.Get("/session/get", sessionDetailsHandler)
	// Login API
	v1.Post("/session/login", sessionLoginHandler)
	// Logout API
	v1.Get("/session/logout", sessionLogoutHandler)

	// user
	// Signup API
	v1.Post("/signup", userSignupHandler)

	// wallet
	// Get Wallet Details
	v1.Get("/wallet", walletDetailsHandler)
	// Deposit to wallet
	v1.Post("/wallet/make", walletDepositHandler)

}
