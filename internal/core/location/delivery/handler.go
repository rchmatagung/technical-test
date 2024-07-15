package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type LocationHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewLocationHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) LocationHandler {
	return LocationHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}