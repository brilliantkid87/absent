package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func ContactRoutes(app *fiber.App) {
	EmpployeeRepo := repositories.NewContactRepository(psql.DB)
	EmpployeeHandler := handlers.NewContactHandler(EmpployeeRepo)

	app.Post("/api/create-contact", middleware.Auth, EmpployeeHandler.CreateContact)
	app.Put("/api/update-contact", middleware.Auth, EmpployeeHandler.UpdateContact)
}
