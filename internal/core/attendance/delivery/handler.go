package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/attendance/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
	"context"
	"fmt"
	"strconv"

	cm "boilerplate/pkg/constants/message"

	"github.com/gofiber/fiber/v2"
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

func (a AttendanceHandler) AttendanceAbsentIn(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	dataReq := new(models.CreateAttendanceAbsentInRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	respData, errData := a.Usecase.Core.Attendance.AbsentIn(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create attendance absent in", "Sukses membuat absen masuk", respData)
}

func (a AttendanceHandler) AttendanceAbsentOut(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	dataReq := new(models.CreateAttendanceAbsentOutRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	respData, errData := a.Usecase.Core.Attendance.AbsentOut(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create attendance absent out", "Sukses membuat absen keluar", respData)
}

func (a AttendanceHandler) GetAllAttendance(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	respData, total, errData := a.Usecase.Core.Attendance.GetAllAttendance(context.Background())
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get all attendance", "Sukses mendapatkan semua data absensi", respData, 1, 0, total)
}

func (a AttendanceHandler) GetAttendanceById(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	attendance_id, err := strconv.Atoi(c.Query("attendance_id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("attendance_id"), cm.ErrValuePageIntInd("attendance_id"), nil)
	}

	respData, errData := a.Usecase.Core.Attendance.GetAttendanceById(context.Background(), int64(attendance_id))

	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get attendance by id", "Sukses mendapatkan data absensi berdasarkan id", respData)
}

func (a AttendanceHandler) GetAttendanceByEmployeeId(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	employee_id, err := strconv.Atoi(c.Query("employee_id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("employee_id"), cm.ErrValuePageIntInd("employee_id"), nil)
	}

	respData, errData := a.Usecase.Core.Attendance.GetAttendanceByEmployeeId(context.Background(), int64(employee_id))

	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get attendance by employee id", "Sukses mendapatkan data absensi berdasarkan id karyawan", respData)
}

func (a AttendanceHandler) Delete(c *fiber.Ctx) error {
	init := exception.InitException(c, a.Conf, a.Log)

	attendance_id, err := strconv.Atoi(c.Query("attendance_id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("attendance_id"), cm.ErrValuePageIntInd("attendance_id"), nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := a.Usecase.Core.Attendance.Delete(context.Background(), int64(attendance_id), username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}
	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success delete attendance", "Sukses menghapus data absensi", nil)
}
