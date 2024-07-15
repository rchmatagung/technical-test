package models

type CreatePositionRequest struct {
	DepartmentId int64  `json:"department_id" validate:"required"`
	PositionName string `json:"position_name" validate:"required"`
}

type UpdatePositionRequest struct {
	PositionId   int64  `json:"position_id" validate:"required"`
	DepartmentId int64  `json:"department_id" validate:"required"`
	PositionName string `json:"position_name" validate:"required"`
}