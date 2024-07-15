package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type PositionRepo struct {
	DBList *db.DatabaseList
}

func NewPositionRepo(dbList *db.DatabaseList) PositionRepo {
	return PositionRepo{
		DBList: dbList,
	}
}