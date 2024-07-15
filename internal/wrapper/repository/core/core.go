package core

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
	department	"boilerplate/internal/core/department/repository"
	position	"boilerplate/internal/core/position/repository"
	location	"boilerplate/internal/core/location/repository"
	employee	"boilerplate/internal/core/employee/repository"
	attendance	"boilerplate/internal/core/attendance/repository"
)

type CoreRepository struct {
	Department	department.Repository
	Position	position.Repository
	Location	location.Repository
	Employee	employee.Repository
	Attendance	attendance.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		Department:	department.NewDepartmentRepo(dbList),
		Position:	position.NewPositionRepo(dbList),
		Location:	location.NewLocationRepo(dbList),
		Employee:	employee.NewEmployeeRepo(dbList),
		Attendance:	attendance.NewAttendanceRepo(dbList),
	}
}
