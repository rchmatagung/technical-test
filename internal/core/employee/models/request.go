package models

type CreateEmployeeRequest struct {
	EmployeeName string `json:"employee_name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	DepartmentId int64  `json:"department_id" validate:"required"`
	PositionId   int64  `json:"position_id" validate:"required"`
	Superior     int64  `json:"superior"`
}

type UpdateEmployeeRequest struct {
	EmployeeId   int64  `json:"employee_id" validate:"required"`
	DepartmentId int64  `json:"department_id" validate:"required"`
	PositionId   int64  `json:"position_id" validate:"required"`
	EmployeeName string `json:"employee_name" validate:"required"`
}

type ForgotPasswordRequest struct {
	EmployeeCode string `json:"employee_code" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type LoginRequest struct {
	EmployeeCode string `json:"employee_code" validate:"required"`
	Password     string `json:"password" validate:"required"`
	RoleName     string `json:"role_name"`
}
