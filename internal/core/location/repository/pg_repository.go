package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type LocationRepo struct {
	DBList *db.DatabaseList
}

func NewLocationRepo(dbList *db.DatabaseList) LocationRepo {
	return LocationRepo{
		DBList: dbList,
	}
}