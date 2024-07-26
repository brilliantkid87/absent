package handlers

import (
	"absent/models"
	"absent/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type AttendanceHandler struct {
	Repo      repositories.AttendanceRepository
	Validator *validator.Validate
}

func NewAttendanceHandler(repo repositories.AttendanceRepository) *AttendanceHandler {
	return &AttendanceHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *AttendanceHandler) CreateAttendance(c *fiber.Ctx) error {
	var attendance models.Attendance
	if err := c.BodyParser(&attendance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	params := map[string]interface{}{
		"employee_id": attendance.EmployeeID,
		"location_id": attendance.LocationID,
		"absent_in":   attendance.AbsentIn,
		"absent_out":  attendance.AbsentOut,
		"created_by":  attendance.CreatedBy,
		"updated_by":  attendance.UpdatedBy,
	}

	attendanceID, err := h.Repo.CreateAttendance(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create attendance"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"attendance_id": attendanceID})
}

func (h *AttendanceHandler) UpdateAttendance(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.UpdateAttendance(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update attendance"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Update master attendance successful"})
	} else {
		return c.JSON(fiber.Map{"message": "Update master attendance failed"})
	}
}

func (h *AttendanceHandler) DeleteAttendance(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.DeleteAttendance(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete attendance"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Delete attendance successfully"})
	}
	return c.JSON(fiber.Map{"message": "Delete attendance failed"})
}

func (h *AttendanceHandler) GetAllAttendance(c *fiber.Ctx) error {
	Employees, err := h.Repo.GetAllAttendance()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve Employees"})
	}

	return c.JSON(fiber.Map{"Employees": Employees})
}

func (h *AttendanceHandler) GetAbsenceReport(c *fiber.Ctx) error {
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	if startTime == "" || endTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "start_time and end_time query parameters are required"})
	}

	params := map[string]string{
		"start_time": startTime,
		"end_time":   endTime,
	}

	reports, err := h.Repo.GetAbsenceReport(params)
	if err != nil {
		log.Println("Failed to get absence report:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get absence report"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"reports": reports})
}
