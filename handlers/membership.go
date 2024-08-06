package handlers

import (
	"absent/models"
	"absent/pkg/bcrypt"
	jwtToken "absent/pkg/jwt"
	"absent/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/go-playground/validator.v9"
)

type MembershipHandler struct {
	Repo      repositories.MembershipRepository
	Validator *validator.Validate
}

func NewMembershipHandler(repo repositories.MembershipRepository) *MembershipHandler {
	return &MembershipHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *MembershipHandler) LoginMembership(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	membershipName, membershipNameExists := params["name"].(string)
	password, passwordExists := params["password"].(string)

	if !membershipNameExists || !passwordExists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "membership name and password are required"})
	}

	membershipParams := map[string]interface{}{
		"name": membershipName,
	}

	memberships, err := h.Repo.GetAllMemberships(membershipParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check employee"})
	}

	if len(memberships) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "membership name not registered. Please register first."})
	}

	membership := memberships[0]

	// Check password
	passwordHash, ok := membership["password"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Password hash is invalid"})
	}

	// Verify password
	err = bcrypt.CheckPasswordHash(password, passwordHash)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid membership name or password"})
	}

	membershipID, ok := membership["membership_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "membership ID is invalid"})
	}

	membershipIDStr := fmt.Sprintf("%d", int(membershipID))

	// Generate JWT token
	claims := jwt.MapClaims{
		"id":  membershipIDStr,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"employee": membership,
		"token":    token,
	})
}

func (h *MembershipHandler) CreateMembership(c *fiber.Ctx) error {
	var membership models.Membership
	if err := c.BodyParser(&membership); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// // Validate request body
	if err := h.Validator.Struct(&membership); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	password, err := bcrypt.HashingPassword(membership.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hashing pass failed"})
	}

	params := map[string]interface{}{
		"name":       membership.Name,
		"password":   password,
		"address":    membership.Address,
		"is_active":  membership.IsActive,
		"created_by": membership.CreatedBy,
	}

	// Call repository to create employee
	membershipID, err := h.Repo.CreateMembership(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"membership_id": membershipID})
}

func (h *MembershipHandler) GetActiveMembershipsWithContacts(c *fiber.Ctx) error {
	memberships, err := h.Repo.GetActiveMembershipsWithContacts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if memberships == nil {
		return c.JSON([]interface{}{})
	}

	return c.JSON(memberships)
}
