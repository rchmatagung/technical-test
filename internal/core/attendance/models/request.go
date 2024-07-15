package models

type CreateAttendanceAbsentInRequest struct {
	EmployeeId int64     `json:"employee_id" validate:"required"`
	LocationId int64     `json:"location_id" validate:"required"`
}

type CreateAttendanceAbsentOutRequest struct {
	EmployeeId int64  `json:"employee_id" validate:"required"`
	LocationId int64  `json:"location_id" validate:"required"`
}

type UpdateAttendanceRequest struct {
	AttendanceId int64  `json:"attendance_id" validate:"required"`
	EmployeeId   int64  `json:"employee_id" validate:"required"`
	LocationId   int64  `json:"location_id" validate:"required"`
	AbsentIn     string `json:"absent_in" validate:"required"`
	AbsentOut    string `json:"absent_out" validate:"required"`
}