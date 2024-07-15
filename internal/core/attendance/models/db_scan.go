package models

import "time"

type GetAttendanceScan struct {
	AttendanceId int64      `json:"attendance_id"`
	EmployeeId   int64      `json:"employee_id"`
	EmployeeName string     `json:"employee_name"`
	LocationId   int64      `json:"location_id"`
	LocationName string     `json:"location_name"`
	AbsentIn     time.Time  `json:"absent_in"`
	AbsentOut    time.Time  `json:"absent_out"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedBy    *string    `json:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type ListAttendanceScan struct {
	AttendanceId int64      `json:"attendance_id"`
	EmployeeId   int64      `json:"employee_id"`
	EmployeeName string     `json:"employee_name"`
	LocationId   int64      `json:"location_id"`
	LocationName string     `json:"location_name"`
	AbsentIn     time.Time  `json:"absent_in"`
	AbsentOut    time.Time  `json:"absent_out"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedBy    *string    `json:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
