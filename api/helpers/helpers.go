package helpers

import (
	"context"

	"github.com/golang-demos/ecommerce-basic/database"
	"github.com/golang-demos/ecommerce-basic/models"
)

func GetUserBySessionId(sessToken string) *models.User {
	collection := database.Db.Collection("users")
	ctx := context.Background()

	var user models.User
	collection.FindOne(ctx, models.User{Token: sessToken}).Decode(&user)

	return &user
}
