package models

type Employee struct {
	EmployeeID   int     `json:"employee_id"`
	EmployeeCode string  `json:"employee_code"`
	EmployeeName string  `json:"employee_name"`
	Password     string  `json:"password"`
	DepartmentID int     `json:"department_id"`
	PositionID   int     `json:"position_id"`
	Superior     int     `json:"superior"`
	CreatedAt    string  `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	UpdatedAt    string  `json:"updated_at"`
	UpdatedBy    string  `json:"updated_by"`
	DeletedAt    *string `json:"deleted_at"`
}
