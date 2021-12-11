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

func ConnectDB() {
	ctx := context.Background()
	dbConnectionURI := os.Getenv("ECOMM_DB_CONN_URI")
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionURI))
	Db = Client.Database(os.Getenv("ECOMM_DB_NAME"))
}

func DisconnectDB() {
	ctx := context.Background()
	if err := Client.Disconnect(ctx); err != nil {
		panic(err)
	}
	log.Println("Database disconnected")
}
