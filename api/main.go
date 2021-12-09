package main

import (
	"log"
	"os"

	"github.com/golang-demos/ecommerce-basic/controllers"
	"github.com/golang-demos/ecommerce-basic/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// dbConnectionURI := os.Getenv("ECOMM_DB_CONN_URI")
	apiServerPort := os.Getenv("ECOMM_APP_PORT")

	database.ConnectDB()

	app := fiber.New()

	controllers.RegisterRoutes(app)

	log.Fatal(app.Listen(apiServerPort))
}
