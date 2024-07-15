package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type DepartmentRepo struct {
	DBList *db.DatabaseList
}

func NewDepartmentRepo(dbList *db.DatabaseList) DepartmentRepo {
	return DepartmentRepo{
		DBList: dbList,
	}
}