package models

type MasterDepartment struct {
	DepartmentID   int     `json:"department_id"`
	DepartmentName string  `json:"department_name"`
	CreatedAt      string  `json:"created_at"`
	CreatedBy      string  `json:"created_by"`
	UpdatedAt      string  `json:"updated_at"`
	UpdatedBy      string  `json:"updated_by"`
	DeletedAt      *string `json:"deleted_at"`
}
