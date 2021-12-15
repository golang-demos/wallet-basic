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
	v1.Get("/wallet/:id", walletDetailsHandler)
	// Deposit to wallet
	v1.Post("/wallet/deposit", walletDepositHandler)

	// product
	// Product API
	v1.Post("/product/list", productListHandler)
	// 	- Create, Update, Delete
	v1.Get("/product/:productId", productDetailsHandler)
	v1.Post("/product", productCreateHandler)
	v1.Put("/product/:productId", productUpdateHandler)
	v1.Delete("/product/:productId", productDeleteHandler)

	// variation
	// Variations API
	v1.Get("/variations/list", variationListHandler)
	// 	- Create, Update, Delete
	v1.Get("/variation/:productId", variationDetailsHandler)
	v1.Post("/variation", variationCreateHandler)
	v1.Put("/variation/:variationId", variationUpdateHandler)
	v1.Delete("/variation/:variationId", variationDeleteHandler)

	// order
	// Order API
	v1.Get("/order/list", orderListHandler)
	// 	- Create, Update, Cancel
	v1.Get("/order/:productId", orderDetailsHandler)
	v1.Post("/order", orderCreateHandler)
	v1.Put("/order/:orderId", orderUpdateHandler)
	v1.Delete("/order/:orderId", orderDeleteHandler)

}
