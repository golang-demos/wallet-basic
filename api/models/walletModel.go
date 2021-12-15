package models

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/golang-demos/ecommerce-basic/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id,"`
	Balance float32            `json:"balance" bson:"balance"`
}

func (u *Wallet) Validate() []*ErrorResp {
	return validateModelsForErrors(validator.New().Struct(u))
}

func checkIfWalletExists(user *User) bool {
	var foundWallet Wallet
	cursor := database.WalletColllection.FindOne(context.Background(), bson.M{"user_id": user.ID})
	cursor.Decode(&foundWallet)
	return !foundWallet.ID.IsZero()
}

func CreateWallet(user *User) bool {
	alreadyExists := checkIfWalletExists(user)
	if alreadyExists {
		return false
	}

	wallet := &Wallet{
		UserID:  user.ID,
		Balance: 0.0,
	}
	result, _ := database.WalletColllection.InsertOne(context.Background(), wallet)
	isCreated := false
	if result != nil {
		isCreated = true
	}
	return isCreated
}
