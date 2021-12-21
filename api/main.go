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
	appMode := "dev"
	argsLength := len(os.Args)
	if argsLength > 0 {
		if os.Args[0] == "prod" {
			appMode = "prod"
		}
	}

	envFile := ".env.dev"
	if appMode == "prod" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
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
