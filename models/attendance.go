package models

type Attendance struct {
	AttendanceID int     `json:"attendance_id"`
	EmployeeID   int     `json:"employee_id"`
	LocationID   int     `json:"location_id"`
	AbsentIn     string  `json:"absent_in"`
	AbsentOut    string  `json:"absent_out"`
	CreatedAt    string  `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	UpdatedAt    string  `json:"updated_at"`
	UpdatedBy    string  `json:"updated_by"`
	DeletedAt    *string `json:"deleted_at"`
}

type ReportAbsence struct {
	AbsentIn       string `json:"absent_in"`
	AbsentOut      string `json:"absent_out"`
	EmployeeCode   string `json:"employee_code"`
	EmployeeName   string `json:"employee_name"`
	DepartmentName string `json:"department_name"`
	PositionName   string `json:"position_name"`
	LocationName   string `json:"location_name"`
}
