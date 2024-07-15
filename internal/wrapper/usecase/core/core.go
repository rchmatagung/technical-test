package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
	department	"boilerplate/internal/core/department/usecase"
	position	"boilerplate/internal/core/position/usecase"
	location	"boilerplate/internal/core/location/usecase"
	employee	"boilerplate/internal/core/employee/usecase"
	attendance	"boilerplate/internal/core/attendance/usecase"
)

type CoreUsecase struct {
	Department	department.Usecase
	Position	position.Usecase
	Location	location.Usecase
	Employee	employee.Usecase
	Attendance	attendance.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		Department:	department.NewDepartmentUsecase(repo, conf, dbList, log),
		Position:	position.NewPositionUsecase(repo, conf, dbList, log),
		Location:	location.NewLocationUsecase(repo, conf, dbList, log),
		Employee:	employee.NewEmployeeUsecase(repo, conf, dbList, log),
		Attendance:	attendance.NewAttendanceUsecase(repo, conf, dbList, log),
	}
}
