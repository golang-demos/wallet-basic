package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
	"go.mongodb.org/mongo-driver/bson"
)

func walletDetailsHandler(c *fiber.Ctx) error {
	userObjectId := c.Locals("SESSION_USER_ID")

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

	userObjectId := c.Locals("SESSION_USER_ID")
	var user models.User
	err := database.UserCollection.FindOne(context.Background(), bson.M{
		"_id": userObjectId,
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
		"user_id": userObjectId,
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

func walletStatementHandler(c *fiber.Ctx) error {
	userObjectId := c.Locals("SESSION_USER_ID")
	var transactionList []fiber.Map
	cursor, err := database.TransactionColllection.Find(context.Background(), bson.M{
		"user_id": userObjectId,
	})
	if err != nil {
		return c.JSON(transactionList)
	}

	ctx := context.Background()
	for cursor.Next(ctx) {
		var transaction models.Transaction
		cursor.Decode(&transaction)
		transactionList = append(transactionList, fiber.Map{
			"amount":     transaction.Amount,
			"type":       transaction.TransType,
			"created_at": time.Unix(int64(transaction.CreatedAt.T), 0).String(),
		})
	}

	return c.JSON(transactionList)
}
