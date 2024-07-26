package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App) {
	EmpployeeRepo := repositories.NewEmployeeRepository(psql.DB)
	EmpployeeHandler := handlers.NewEmployeeHandler(EmpployeeRepo)

	app.Post("/api/create-employee", EmpployeeHandler.CreateEmployee)
	app.Post("/api/update-employee", middleware.Auth, EmpployeeHandler.UpdateEmployee)
	app.Post("/api/delete-employee", middleware.Auth, EmpployeeHandler.DeleteEmployee)
	app.Get("/api/getall-employee", middleware.Auth, EmpployeeHandler.GetAllEmployee)
	app.Post("/api/login", EmpployeeHandler.LoginEmployee)
}
