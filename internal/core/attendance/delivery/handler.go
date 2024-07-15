package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type AttendanceHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewAttendanceHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) AttendanceHandler {
	return AttendanceHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}