package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func MasterLocRoutes(app *fiber.App) {
	MasterLocRepo := repositories.NewLocationRepository(psql.DB)
	MasterLocHandler := handlers.NewLocationHandler(MasterLocRepo)

	app.Post("/api/create-masterloc", middleware.Auth, MasterLocHandler.CreateLocation)
	app.Post("/api/update-masterloc", middleware.Auth, MasterLocHandler.UpdateLocation)
	app.Post("/api/delete-masterloc", middleware.Auth, MasterLocHandler.DeleteLocation)
	app.Get("/api/getall-masterloc", middleware.Auth, MasterLocHandler.GetAllLocations)
}
