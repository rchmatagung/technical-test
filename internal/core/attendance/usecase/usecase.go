package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
}

type AttendanceUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewAttendanceUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) AttendanceUsecase {
	return AttendanceUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}