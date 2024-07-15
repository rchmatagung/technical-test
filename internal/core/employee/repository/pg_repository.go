package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type EmployeeRepo struct {
	DBList *db.DatabaseList
}

func NewEmployeeRepo(dbList *db.DatabaseList) EmployeeRepo {
	return EmployeeRepo{
		DBList: dbList,
	}
}