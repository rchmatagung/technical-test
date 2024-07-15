package models

import "time"

type GetPosition struct {
	PositionId     int64      `json:"position_id"`
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	PositionName   string     `json:"position_name"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type ListPosition struct {
	PositionId     int64      `json:"position_id"`
	DepartmentId   int64      `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	PositionName   string     `json:"position_name"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      *string    `json:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
