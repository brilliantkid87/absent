package repositories

import (
	"absent/models"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(params map[string]interface{}) (int, error)
	UpdateEmployee(params map[string]interface{}) (bool, error)
	DeleteEmployee(params map[string]interface{}) (bool, error)
	GetAllEmployees(params map[string]interface{}) ([]map[string]interface{}, error)
	GetAllEmployee() ([]models.Employee, error)
}

func NewEmployeeRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateEmployee(params map[string]interface{}) (int, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var employeeID int
	query := `SELECT employee.create_employee($1::jsonb)`

	if err := r.db.Raw(query, string(paramsJSON)).Scan(&employeeID).Error; err != nil {
		if strings.Contains(err.Error(), "Employee Code") && strings.Contains(err.Error(), "already exists") {
			return 0, fmt.Errorf("employee code %s already exists", params["employee_code"])
		}

		return 0, err
	}

	return employeeID, nil
}

func (r *repository) UpdateEmployee(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT employee.update_employee($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) DeleteEmployee(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT employee.delete_employee($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) GetAllEmployees(params map[string]interface{}) ([]map[string]interface{}, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling:", err)
		return nil, err
	}

	var result []byte
	query := `SELECT employee.get_all_employees($1::jsonb)`

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

func (r *repository) GetAllEmployee() ([]models.Employee, error) {
	var result []byte
	query := `SELECT employee.getall_employee()`

	row := r.db.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing getall_master_location:", err)
		return nil, err
	}

	var employees []models.Employee
	if err := json.Unmarshal(result, &employees); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	const timeLayout = "2006-01-02T15:04:05.999999"

	// Parse time strings using the custom layout
	for i := range employees {
		if employees[i].CreatedAt != "" {
			if _, err := time.Parse(timeLayout, employees[i].CreatedAt); err != nil {
				log.Printf("Error parsing CreatedAt for employee %d: %v", employees[i].EmployeeID, err)
			}
		}
		if employees[i].UpdatedAt != "" {
			if _, err := time.Parse(timeLayout, employees[i].UpdatedAt); err != nil {
				log.Printf("Error parsing UpdatedAt for employee %d: %v", employees[i].EmployeeID, err)
			}
		}
		if employees[i].DeletedAt != nil {
			if _, err := time.Parse(timeLayout, *employees[i].DeletedAt); err != nil {
				log.Printf("Error parsing DeletedAt for employee %d: %v", employees[i].EmployeeID, err)
			}
		}
	}

	return employees, nil
}
