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
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	defer database.DisconnectDB()

	app := fiber.New()

	controllers.RegisterRoutes(app)

	apiServerPort := os.Getenv("ECOMM_APP_PORT")
	log.Fatal(app.Listen(apiServerPort))
}
