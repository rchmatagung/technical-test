package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
}

type PositionUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewPositionUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) PositionUsecase {
	return PositionUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}