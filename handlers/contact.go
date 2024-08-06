package handlers

import (
	"absent/models"
	"absent/repositories"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type ContactHandler struct {
	Repo      repositories.ContactRepository
	Validator *validator.Validate
}

func NewContactHandler(repo repositories.ContactRepository) *ContactHandler {
	return &ContactHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var contact models.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate request body
	if err := h.Validator.Struct(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	params := map[string]interface{}{
		"membership_id": contact.MembershipID,
		"contact_type":  contact.ContactType,
		"contact_value": contact.ContactValue,
		"is_active":     contact.IsActive,
		"created_by":    contact.CreatedBy,
	}

	// Call repository to create contact
	contactID, err := h.Repo.CreateContact(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contact"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"contact_id": contactID})
}

func (h *ContactHandler) UpdateContact(c *fiber.Ctx) error {
	var contact models.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate request body
	if err := h.Validator.Struct(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	params := map[string]interface{}{
		"contact_id":    contact.ContactID,
		"contact_type":  contact.ContactType,
		"contact_value": contact.ContactValue,
		"is_active":     contact.IsActive,
		"update_by":     contact.UpdateBy,
	}

	// Log parameters for debugging
	fmt.Printf("UpdateContact params: %+v\n", params)

	// Call repository to update contact
	if err := h.Repo.UpdateContact(params); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update contact", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Contact updated successfully"})
}
