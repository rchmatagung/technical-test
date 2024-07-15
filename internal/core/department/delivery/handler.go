package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type DepartmentHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewDepartmentHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) DepartmentHandler {
	return DepartmentHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}