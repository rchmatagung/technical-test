package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
}

type EmployeeUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewEmployeeUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) EmployeeUsecase {
	return EmployeeUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}