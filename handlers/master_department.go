package handlers

import (
	"absent/models"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type DepartmentHandler struct {
	Repo      repositories.MasterDeptRepository
	Validator *validator.Validate
}

func NewDepartmentHandler(repo repositories.MasterDeptRepository) *DepartmentHandler {
	return &DepartmentHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {
	var department models.MasterDepartment

	if err := c.BodyParser(&department); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Create department using repository
	params := map[string]interface{}{
		"department_name": department.DepartmentName,
		"created_by":      department.CreatedBy,
		"updated_by":      department.UpdatedBy,
	}

	departmentID, err := h.Repo.CreateDepartment(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create department"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"department_id": departmentID})
}

func (h *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.UpdateDepartment(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update department"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Update master department successful"})
	} else {
		return c.JSON(fiber.Map{"message": "Update master department failed"})
	}
}

func (h *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.DeleteDepartment(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete department"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Delete master department successful"})
	} else {
		return c.JSON(fiber.Map{"message": "Delete master department failed"})
	}
}

func (h *DepartmentHandler) GetAlldepartments(c *fiber.Ctx) error {
	Departments, err := h.Repo.GetAllDepartment()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve Departments"})
	}

	return c.JSON(fiber.Map{"Departments": Departments})
}
