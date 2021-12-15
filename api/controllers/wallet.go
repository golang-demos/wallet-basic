package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func walletDetailsHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	userObjectId, _ := primitive.ObjectIDFromHex(id)

	var wallet models.Wallet
	database.WalletColllection.FindOne(context.Background(), bson.M{
		"user_id": userObjectId,
	}).Decode(&wallet)

	if wallet.ID.IsZero() {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
			"success": false,
		})
	}

	return c.JSON(wallet)
}

func walletDepositHandler(c *fiber.Ctx) error {
	return nil
}
