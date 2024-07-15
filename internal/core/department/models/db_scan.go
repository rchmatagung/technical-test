package models

import "time"

type GetDepartment struct {
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type ListDepartment struct {
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
