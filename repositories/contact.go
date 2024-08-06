package repositories

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type ContactRepository interface {
	CreateContact(params map[string]interface{}) (int, error)
	UpdateContact(params map[string]interface{}) error
}

func NewContactRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateContact(params map[string]interface{}) (int, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var contactID int
	query := `SELECT public.create_contact($1::jsonb)`

	if err := r.db.Raw(query, string(paramsJSON)).Scan(&contactID).Error; err != nil {
		return 0, err
	}

	return contactID, nil
}

func (r *repository) UpdateContact(params map[string]interface{}) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshalling params: %w", err)
	}

	query := `SELECT public.update_contact($1::jsonb)`

	if err := r.db.Raw(query, string(paramsJSON)).Error; err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}
