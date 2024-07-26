package repositories

import (
	"absent/models"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type MasterDeptRepository interface {
	CreateDepartment(params map[string]interface{}) (string, error)
	// UpdateDepartment(department models.MasterDepartment) (models.MasterDepartment, error)
	UpdateDepartment(params map[string]interface{}) (bool, error)
	DeleteDepartment(params map[string]interface{}) (bool, error)
	GetAllDepartment() ([]models.MasterDepartment, error)
}

func RepositoryMasterDept(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateDepartment(params map[string]interface{}) (string, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	query := `SELECT master_department.create_master_department($1::jsonb)`

	var departmentID string
	if err := r.db.Raw(query, string(paramsJSON)).Scan(&departmentID).Error; err != nil {
		return "", err
	}

	return departmentID, nil
}

func (r *repository) UpdateDepartment(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_department.update_master_department($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *repository) DeleteDepartment(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT master_department.delete_master_department($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *repository) GetAllDepartment() ([]models.MasterDepartment, error) {
	var result []byte
	query := `SELECT master_department.getall_master_department()`

	row := r.db.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing getall_master_department:", err)
		return nil, err
	}

	var departments []models.MasterDepartment
	if err := json.Unmarshal(result, &departments); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	const timeLayout = "2006-01-02T15:04:05.999999"

	// Parse time strings using the custom layout
	for i := range departments {
		if departments[i].CreatedAt != "" {
			if _, err := time.Parse(timeLayout, departments[i].CreatedAt); err != nil {
				log.Printf("Error parsing CreatedAt for location %d: %v", departments[i].DepartmentID, err)
			}
		}
		if departments[i].UpdatedAt != "" {
			if _, err := time.Parse(timeLayout, departments[i].UpdatedAt); err != nil {
				log.Printf("Error parsing UpdatedAt for location %d: %v", departments[i].DepartmentID, err)
			}
		}
		if departments[i].DeletedAt != nil {
			if _, err := time.Parse(timeLayout, *departments[i].DeletedAt); err != nil {
				log.Printf("Error parsing DeletedAt for location %d: %v", departments[i].DepartmentID, err)
			}
		}
	}

	return departments, nil
}
