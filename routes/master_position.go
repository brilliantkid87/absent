package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func MasterPostRoutes(app *fiber.App) {
	MasterPostRepo := repositories.NewPositionRepository(psql.DB)
	MasterPostHandler := handlers.NewPositionHandler(MasterPostRepo)

	app.Post("/api/create-masterpost", middleware.Auth, MasterPostHandler.CreatePosition)
	app.Post("/api/update-masterpost", middleware.Auth, MasterPostHandler.UpdatePosition)
	app.Post("/api/delete-masterpost", middleware.Auth, MasterPostHandler.DeletePosition)
	app.Get("/api/getall-masterpost", middleware.Auth, MasterPostHandler.GetAllPositions)
}
