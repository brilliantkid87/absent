package routes

import (
	"absent/handlers"
	"absent/pkg/middleware"
	"absent/pkg/psql"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(app *fiber.App) {
	attendanceRepo := repositories.NewAttendanceRepository(psql.DB)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceRepo)

	app.Post("/api/create-attendance", middleware.Auth, attendanceHandler.CreateAttendance)
	app.Post("/api/update-attendance", middleware.Auth, attendanceHandler.UpdateAttendance)
	app.Post("/api/delete-attendance", middleware.Auth, attendanceHandler.DeleteAttendance)
	app.Get("/api/getall-attendance", middleware.Auth, attendanceHandler.GetAllAttendance)
	app.Get("/api/report-attendance", middleware.Auth, attendanceHandler.GetAbsenceReport)
}
