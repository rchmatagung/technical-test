package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type PositionHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewPositionHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) PositionHandler {
	return PositionHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}