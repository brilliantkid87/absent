package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	CreateMembership(params map[string]interface{}) (int, error)
	GetAllMemberships(params map[string]interface{}) ([]map[string]interface{}, error)
	GetActiveMembershipsWithContacts() (json.RawMessage, error)
}

func NewMembershipRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMembership(params map[string]interface{}) (int, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var membershipID int
	query := `SELECT public.membership($1::jsonb)`

	if err := r.db.Raw(query, string(paramsJSON)).Scan(&membershipID).Error; err != nil {
		return 0, err
	}

	return membershipID, nil
}

func (r *repository) GetAllMemberships(params map[string]interface{}) ([]map[string]interface{}, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling:", err)
		return nil, err
	}

	var result []byte
	query := `SELECT public.getall_membership($1::jsonb)`

	row := r.db.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing get_all_users:", err)
		return nil, err
	}

	var employees []map[string]interface{}
	if err := json.Unmarshal(result, &employees); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	return employees, nil
}

func (r *repository) GetActiveMembershipsWithContacts() (json.RawMessage, error) {
	var result []byte
	query := `SELECT public.get_active_memberships_with_contacts()`

	// Execute raw SQL query
	err := r.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	// Convert []byte to json.RawMessage
	return json.RawMessage(result), nil
}
