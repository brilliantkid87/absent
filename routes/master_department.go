package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func MasterDeptRoutes(app *fiber.App) {
	MasterDeptRepo := repositories.RepositoryMasterDept(psql.DB)
	MasterDeptHandler := handlers.NewDepartmentHandler(MasterDeptRepo)

	app.Post("/api/create-masterdept", middleware.Auth, MasterDeptHandler.CreateDepartment)
	app.Post("/api/update-masterdept", middleware.Auth, MasterDeptHandler.UpdateDepartment)
	app.Post("/api/delete-masterdept", middleware.Auth, MasterDeptHandler.DeleteDepartment)
	app.Get("/api/getall-masterdept", middleware.Auth, MasterDeptHandler.GetAlldepartments)
}
