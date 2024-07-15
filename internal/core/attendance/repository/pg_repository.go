package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type AttendanceRepo struct {
	DBList *db.DatabaseList
}

func NewAttendanceRepo(dbList *db.DatabaseList) AttendanceRepo {
	return AttendanceRepo{
		DBList: dbList,
	}
}