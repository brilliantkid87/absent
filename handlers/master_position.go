package handlers

import (
	"absent/models"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type PositionHandler struct {
	Repo      repositories.PositionRepository
	Validator *validator.Validate
}

func NewPositionHandler(repo repositories.PositionRepository) *PositionHandler {
	return &PositionHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *PositionHandler) CreatePosition(c *fiber.Ctx) error {
	var position models.Position
	if err := c.BodyParser(&position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	params := map[string]interface{}{
		"department_id": position.DepartmentID,
		"position_name": position.PositionName,
		"created_by":    position.CreatedBy,
	}

	positionID, err := h.Repo.CreatePosition(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create position"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"position_id": positionID})
}

func (h *PositionHandler) UpdatePosition(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.UpdatePosition(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update position"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Update master position successful"})
	}
	return c.JSON(fiber.Map{"message": "Update master position failed"})
}

func (h *PositionHandler) DeletePosition(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.DeletePosition(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete position"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Position successfully deleted"})
	}
	return c.JSON(fiber.Map{"message": "Position deletion failed"})
}

func (h *PositionHandler) GetAllPositions(c *fiber.Ctx) error {
	positions, err := h.Repo.GetAllPositions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve positions"})
	}

	return c.JSON(fiber.Map{"positions": positions})
}
