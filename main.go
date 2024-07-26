package main

import (
	"absent/pkg/psql"
	"absent/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,DELETE",
		AllowHeaders: "X-Requested-With, Content-Type, Authorization",
	}))

	psql.DatabaseConnection()

	routes.RouteInit(app)

	app.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")

	// Start the server
	fmt.Println("Server running on localhost:" + port)
	err := app.Listen(":" + port)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
