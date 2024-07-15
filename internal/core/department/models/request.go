package models

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" validate:"required"`
}

type UpdateDepartmentRequest struct {
	DepartmentId   int64  `json:"department_id" validate:"required"`
	DepartmentName string `json:"department_name" validate:"required"`
}