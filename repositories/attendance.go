package repositories

import (
	"absent/models"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	CreateAttendance(params map[string]interface{}) (int, error)
	UpdateAttendance(params map[string]interface{}) (bool, error)
	DeleteAttendance(params map[string]interface{}) (bool, error)
	GetAllAttendance() ([]models.Attendance, error)
	GetAbsenceReport(params map[string]string) ([]models.ReportAbsence, error)
}

func NewAttendanceRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAttendance(params map[string]interface{}) (int, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	var attendanceID int
	err = r.db.Raw("SELECT attendance.create_attendance($1::jsonb)", string(jsonData)).Scan(&attendanceID).Error
	if err != nil {
		return 0, err
	}

	return attendanceID, nil
}

func (r *repository) UpdateAttendance(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT attendance.update_attendance($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) DeleteAttendance(params map[string]interface{}) (bool, error) {
	paramsJSONB, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	var result bool

	err = r.db.Raw("SELECT attendance.delete_attendance($1::jsonb)", paramsJSONB).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) GetAllAttendance() ([]models.Attendance, error) {
	var result []byte
	query := `SELECT attendance.getall_attendance()`

	row := r.db.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing getall_master_location:", err)
		return nil, err
	}

	var attendances []models.Attendance
	if err := json.Unmarshal(result, &attendances); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	const timeLayout = "2006-01-02T15:04:05.999999"

	// Parse time strings using the custom layout
	for i := range attendances {
		if attendances[i].CreatedAt != "" {
			if _, err := time.Parse(timeLayout, attendances[i].CreatedAt); err != nil {
				log.Printf("Error parsing CreatedAt for employee %d: %v", attendances[i].AttendanceID, err)
			}
		}
		if attendances[i].UpdatedAt != "" {
			if _, err := time.Parse(timeLayout, attendances[i].UpdatedAt); err != nil {
				log.Printf("Error parsing UpdatedAt for employee %d: %v", attendances[i].AttendanceID, err)
			}
		}
		if attendances[i].DeletedAt != nil {
			if _, err := time.Parse(timeLayout, *attendances[i].DeletedAt); err != nil {
				log.Printf("Error parsing DeletedAt for employee %d: %v", attendances[i].AttendanceID, err)
			}
		}
	}

	return attendances, nil
}

func (r *repository) GetAbsenceReport(params map[string]string) ([]models.ReportAbsence, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling params:", err)
		return nil, err
	}

	var resultJSON []byte
	query := `SELECT attendance.report_absence($1::jsonb)`
	row := r.db.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&resultJSON); err != nil {
		log.Println("Error executing report_absence:", err)
		return nil, err
	}

	var reports []models.ReportAbsence
	if err := json.Unmarshal(resultJSON, &reports); err != nil {
		log.Println("Error unmarshaling result:", err)
		return nil, err
	}

	return reports, nil
}
