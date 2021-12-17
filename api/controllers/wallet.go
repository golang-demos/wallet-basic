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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"balance": wallet.Balance,
	})
}

func walletDepositHandler(c *fiber.Ctx) error {
	type makeTransactionData struct {
		TransType string  `json:"trans_type"`
		Amount    float32 `json:"amount"`
	}

	var postData makeTransactionData
	c.BodyParser(&postData)

	UserId := c.Params("id")
	UserObjectId, _ := primitive.ObjectIDFromHex(UserId)
	var user models.User
	err := database.UserCollection.FindOne(context.Background(), bson.M{
		"_id": UserObjectId,
	}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Not Found",
		})
	}

	if postData.TransType != "credit" && postData.TransType != "debit" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid Transaction type",
		})
	}

	var wallet models.Wallet
	err = database.WalletColllection.FindOne(context.Background(), bson.M{
		"user_id": UserObjectId,
	}).Decode(&wallet)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Wallet not found",
		})
	}

	transaction, ok := models.MakeTransaction(wallet, postData.TransType, postData.Amount)
	if !ok || transaction.ID.IsZero() {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create transaction",
		})
	}

	if postData.TransType == "credit" {
		wallet.Balance = wallet.Balance + transaction.Amount
	} else if postData.TransType == "debit" {
		wallet.Balance = wallet.Balance - transaction.Amount
	}
	filter := bson.M{"_id": wallet.ID}
	udpateData := bson.D{{"$set", bson.D{{"balance", wallet.Balance}}}}
	_, err = database.WalletColllection.UpdateOne(context.Background(), filter, udpateData)

	if err != nil {
		// Delete transaction at this point
		database.TransactionColllection.DeleteOne(context.Background(), bson.M{"_id": transaction.ID})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to update balance",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"balance": wallet.Balance,
	})
}
