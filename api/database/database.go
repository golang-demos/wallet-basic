package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Db *mongo.Database

var UserCollection *mongo.Collection
var WalletColllection *mongo.Collection
var TransactionColllection *mongo.Collection
var ProductColllection *mongo.Collection
var VariationColllection *mongo.Collection
var OrderColllection *mongo.Collection

func ConnectDB() {
	dbConnectionURI := os.Getenv("ECOMM_DB_CONN_URI")
	Client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbConnectionURI))
	if err != nil {
		panic(err)
	}
	Db = Client.Database(os.Getenv("ECOMM_DB_NAME"))
	UserCollection = Db.Collection("users")
	WalletColllection = Db.Collection("wallets")
	TransactionColllection = Db.Collection("transactions")
	ProductColllection = Db.Collection("products")
	VariationColllection = Db.Collection("variations")
	OrderColllection = Db.Collection("orders")
}

func DisconnectDB() {
	if err := Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Database disconnected")
}
