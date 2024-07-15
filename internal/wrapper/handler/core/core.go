package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
	department	"boilerplate/internal/core/department/delivery"
	position	"boilerplate/internal/core/position/delivery"
	location	"boilerplate/internal/core/location/delivery"
	employee	"boilerplate/internal/core/employee/delivery"
	attendance	"boilerplate/internal/core/attendance/delivery"
)

type CoreHandler struct {
	Department	department.DepartmentHandler
	Position	position.PositionHandler
	Location	location.LocationHandler
	Employee	employee.EmployeeHandler
	Attendance	attendance.AttendanceHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		Department:	department.NewDepartmentHandler(uc, conf, log),
		Position:	position.NewPositionHandler(uc, conf, log),
		Location:	location.NewLocationHandler(uc, conf, log),
		Employee:	employee.NewEmployeeHandler(uc, conf, log),
		Attendance:	attendance.NewAttendanceHandler(uc, conf, log),
	}
}
