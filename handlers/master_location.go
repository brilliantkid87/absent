package handlers

import (
	"absent/models"
	"absent/repositories"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type LocationHandler struct {
	Repo      repositories.LocationRepository
	Validator *validator.Validate
}

func NewLocationHandler(repo repositories.LocationRepository) *LocationHandler {
	return &LocationHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *LocationHandler) CreateLocation(c *fiber.Ctx) error {
	var location models.Location
	if err := c.BodyParser(&location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	params := map[string]interface{}{
		"location_name": location.LocationName,
		"created_by":    location.CreatedBy,
	}

	locationID, err := h.Repo.CreateLocation(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create location"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"location_id": locationID})
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.UpdateLocation(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update position"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Update master position successful"})
	}
	return c.JSON(fiber.Map{"message": "Update master position failed"})
}

func (h *LocationHandler) DeleteLocation(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.DeleteLocation(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete Location"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Location successfully deleted"})
	}
	return c.JSON(fiber.Map{"message": "Location deletion failed"})
}

func (h *LocationHandler) GetAllLocations(c *fiber.Ctx) error {
	Locations, err := h.Repo.GetAllLocation()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve Locations"})
	}

	return c.JSON(fiber.Map{"Locations": Locations})
}
