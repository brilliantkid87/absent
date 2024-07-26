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

type EmployeeHandler struct {
	Repo      repositories.EmployeeRepository
	Validator *validator.Validate
}

func NewEmployeeHandler(repo repositories.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// // Validate request body
	if err := h.Validator.Struct(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	password, err := bcrypt.HashingPassword(employee.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hashing pass failed"})
	}

	employeeParams := map[string]interface{}{
		"employee_name": employee.EmployeeName,
	}

	employees, err := h.Repo.GetAllEmployees(employeeParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check employee"})
	}

	if len(employees) > 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Employee name is already registered"})
	}
	// Prepare parameters for the repository
	params := map[string]interface{}{
		"employee_name": employee.EmployeeName,
		"password":      password,
		"employee_code": employee.EmployeeCode,
		"department_id": employee.DepartmentID,
		"position_id":   employee.PositionID,
		"superior":      employee.Superior,
		"created_by":    employee.CreatedBy,
	}

	// Call repository to create employee
	employeeID, err := h.Repo.CreateEmployee(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"employee_id": employeeID})
}

func (h *EmployeeHandler) LoginEmployee(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	employeeName, employeeNameExists := params["employee_name"].(string)
	password, passwordExists := params["password"].(string)

	if !employeeNameExists || !passwordExists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Employee name and password are required"})
	}

	employeeParams := map[string]interface{}{
		"employee_name": employeeName,
	}

	employees, err := h.Repo.GetAllEmployees(employeeParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check employee"})
	}

	if len(employees) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Employee name not registered. Please register first."})
	}

	employee := employees[0]

	// Check password
	passwordHash, ok := employee["password"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Password hash is invalid"})
	}

	// Verify password
	err = bcrypt.CheckPasswordHash(password, passwordHash)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid employee name or password"})
	}

	employeeID, ok := employee["employee_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Employee ID is invalid"})
	}

	employeeIDStr := fmt.Sprintf("%d", int(employeeID))

	// Generate JWT token
	claims := jwt.MapClaims{
		"id":  employeeIDStr,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"employee": employee,
		"token":    token,
	})
}

func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate request body
	if err := h.Validator.Struct(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
	}

	// Hash password
	password, err := bcrypt.HashingPassword(employee.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Hashing password failed"})
	}

	// Prepare parameters for the repository
	params := map[string]interface{}{
		"employee_id":   employee.EmployeeID,
		"employee_code": employee.EmployeeCode,
		"employee_name": employee.EmployeeName,
		"password":      password,
		"department_id": employee.DepartmentID,
		"position_id":   employee.PositionID,
		"superior":      employee.Superior,
		"created_by":    employee.CreatedBy,
		"updated_by":    employee.UpdatedBy,
		"deleted_at":    employee.DeletedAt,
	}

	// Call the repository method
	success, err := h.Repo.UpdateEmployee(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update employee"})
	}

	// Respond based on the result
	if success {
		return c.JSON(fiber.Map{"message": "Update employee successful"})
	}
	return c.JSON(fiber.Map{"message": "Update employee failed"})
}

func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	var params map[string]interface{}
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	success, err := h.Repo.DeleteEmployee(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete EmDeleteEmployee"})
	}

	if success {
		return c.JSON(fiber.Map{"message": "Delete Employee successfully"})
	}
	return c.JSON(fiber.Map{"message": "Delete Employee failed"})
}

func (h *EmployeeHandler) GetAllEmployee(c *fiber.Ctx) error {
	Employees, err := h.Repo.GetAllEmployee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve Employees"})
	}

	return c.JSON(fiber.Map{"Employees": Employees})
}
