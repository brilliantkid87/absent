package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func MembershipRoutes(app *fiber.App) {
	EmpployeeRepo := repositories.NewMembershipRepository(psql.DB)
	EmpployeeHandler := handlers.NewMembershipHandler(EmpployeeRepo)

	app.Post("/api/login-membership", EmpployeeHandler.LoginMembership)
	app.Post("/api/create-membership", middleware.Auth, EmpployeeHandler.CreateMembership)
	app.Get("/api/getall-membership", middleware.Auth, EmpployeeHandler.GetActiveMembershipsWithContacts)
}
