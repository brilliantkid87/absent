package models

type Location struct {
	LocationID   int     `json:"location_id"`
	LocationName string  `json:"location_name"`
	CreatedAt    string  `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	UpdatedAt    string  `json:"updated_at"`
	UpdatedBy    string  `json:"updated_by"`
	DeletedAt    *string `json:"deleted_at"`
}
