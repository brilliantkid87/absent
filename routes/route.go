package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	MasterDeptRoutes(app)
	EmployeeRoutes(app)
	MasterLocRoutes(app)
	MasterPostRoutes(app)
	AttendanceRoutes(app)
	MembershipRoutes(app)
	ContactRoutes(app)
}
