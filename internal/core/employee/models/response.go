package models

type LoginResponse struct {
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
	AccessToken  string `json:"access_token"`
	RoleName     string `json:"role_name"`
}