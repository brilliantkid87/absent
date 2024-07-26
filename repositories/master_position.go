package repositories

import (
	"absent/models"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type PositionRepository interface {
	CreatePosition(params map[string]interface{}) (int, error)
	UpdatePosition(params map[string]interface{}) (bool, error)
	DeletePosition(params map[string]interface{}) (bool, error)
	GetAllPositions() ([]models.Position, error)
}

func NewPositionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePosition(params map[string]interface{}) (int, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var positionID int
	err = r.db.Raw("SELECT master_position.create_master_position($1::jsonb)", string(jsonData)).Scan(&positionID).Error
	if err != nil {
		return 0, err
	}

	return positionID, nil
}

func (r *repository) UpdatePosition(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_position.update_master_position($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) DeletePosition(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_position.delete_master_position($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) GetAllPositions() ([]models.Position, error) {
	var result []byte
	query := `SELECT master_position.getall_master_position()`

	row := r.db.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing getall_master_position:", err)
		return nil, err
	}

	var positions []models.Position
	if err := json.Unmarshal(result, &positions); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	const timeLayout = "2006-01-02T15:04:05.999999"

	// Parse time strings using the custom layout
	for i := range positions {
		if positions[i].CreatedAt != "" {
			if _, err := time.Parse(timeLayout, positions[i].CreatedAt); err != nil {
				log.Printf("Error parsing CreatedAt for position %d: %v", positions[i].PositionID, err)
			}
		}
		if positions[i].UpdatedAt != "" {
			if _, err := time.Parse(timeLayout, positions[i].UpdatedAt); err != nil {
				log.Printf("Error parsing UpdatedAt for position %d: %v", positions[i].PositionID, err)
			}
		}
		if positions[i].DeletedAt != nil {
			if _, err := time.Parse(timeLayout, *positions[i].DeletedAt); err != nil {
				log.Printf("Error parsing DeletedAt for position %d: %v", positions[i].PositionID, err)
			}
		}
	}

	return positions, nil
}
