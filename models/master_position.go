package models

type Position struct {
	PositionID   int     `json:"position_id"`
	DepartmentID int     `json:"department_id"`
	PositionName string  `json:"position_name"`
	CreatedAt    string  `json:"created_at"` // Use string to handle timestamp format
	CreatedBy    string  `json:"created_by"`
	UpdatedAt    string  `json:"updated_at"` // Use string to handle timestamp format
	UpdatedBy    string  `json:"updated_by"`
	DeletedAt    *string `json:"deleted_at"` // Use pointer to handle NULL values
}
