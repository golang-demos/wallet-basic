package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	WalletID  primitive.ObjectID `json:"wallet_id,omitempty" bson:"wallet_id,omitempty"`
	Amount    float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	TransType string             `json:"trans_type,omitempty" bson:"trans_type,omitempty"`
}
