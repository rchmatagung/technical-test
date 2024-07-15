package models

import "time"

type GetEmployeeScan struct {
	EmployeeId     int64      `json:"employee_id"`
	EmployeeCode   string     `json:"employee_code"`
	EmployeeName   string     `json:"employee_name"`
	Password       string     `json:"password"`
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	PositionId     int64      `json:"position_id"`
	PositionName   string     `json:"position_name"`
	Superior       int64      `json:"superior"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type ListEmployeeScan struct {
	EmployeeId     int64      `json:"employee_id"`
	EmployeeCode   string     `json:"employee_code"`
	EmployeeName   string     `json:"employee_name"`
	Password       string     `json:"password"`
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	PositionId     int64      `json:"position_id"`
	PositionName   string     `json:"position_name"`
	Superior       int64      `json:"superior"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
