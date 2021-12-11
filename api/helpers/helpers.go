package helpers

import (
	"context"
	"os"

	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
)

func GetUserBySessionId(sessToken string) *models.User {
	database.ConnectDB()
	defer database.DisconnectDB()

	collection := database.Client.Database(os.Getenv("ECOMM_DB_NAME")).Collection("users")
	ctx := context.Background()

	var user models.User
	collection.FindOne(ctx, models.User{Token: sessToken}).Decode(&user)

	return &user
}
