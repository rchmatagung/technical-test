package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type EmployeeHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewEmployeeHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) EmployeeHandler {
	return EmployeeHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}