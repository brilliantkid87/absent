package repositories

import (
	"absent/models"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(params map[string]interface{}) (int, error)
	UpdateLocation(params map[string]interface{}) (bool, error)
	DeleteLocation(params map[string]interface{}) (bool, error)
	GetAllLocation() ([]models.Location, error)
}

func NewLocationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateLocation(params map[string]interface{}) (int, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var locationID int
	err = r.db.Raw("SELECT master_location.create_master_location($1::jsonb)", string(jsonData)).Scan(&locationID).Error
	if err != nil {
		return 0, err
	}

	return locationID, nil
}

func (r *repository) UpdateLocation(params map[string]interface{}) (bool, error) {
	// Convert params to JSONB format
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_location.update_master_location($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) DeleteLocation(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_location.delete_master_location($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) GetAllLocation() ([]models.Location, error) {
	var result []byte
	query := `SELECT master_location.getall_master_location()`

	row := r.db.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing getall_master_location:", err)
		return nil, err
	}

	var locations []models.Location
	if err := json.Unmarshal(result, &locations); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	const timeLayout = "2006-01-02T15:04:05.999999"

	// Parse time strings using the custom layout
	for i := range locations {
		if locations[i].CreatedAt != "" {
			if _, err := time.Parse(timeLayout, locations[i].CreatedAt); err != nil {
				log.Printf("Error parsing CreatedAt for location %d: %v", locations[i].LocationID, err)
			}
		}
		if locations[i].UpdatedAt != "" {
			if _, err := time.Parse(timeLayout, locations[i].UpdatedAt); err != nil {
				log.Printf("Error parsing UpdatedAt for location %d: %v", locations[i].LocationID, err)
			}
		}
		if locations[i].DeletedAt != nil {
			if _, err := time.Parse(timeLayout, *locations[i].DeletedAt); err != nil {
				log.Printf("Error parsing DeletedAt for location %d: %v", locations[i].LocationID, err)
			}
		}
	}

	return locations, nil
}
