package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/employee/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	cm "boilerplate/pkg/constants/message"
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

func (e EmployeeHandler) Create(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	dataReq := new(models.CreateEmployeeRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	respData, errData := e.Usecase.Core.Employee.Create(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create employee", "Sukses membuat karyawan", respData)
}

func (e EmployeeHandler) GetAllEmployee(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	respData, total, errData := e.Usecase.Core.Employee.GetAllEmployee(context.Background())
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get all employee", "Sukses mendapatkan semua karyawan", respData, 1, 0, total)
}

func (e EmployeeHandler) GetEmployeeById(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	employeeId, err := strconv.Atoi(c.Query("employeeId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("employeeId"), cm.ErrValuePageIntInd("employeeId"), nil)
	}

	respData, errData := e.Usecase.Core.Employee.GetEmployeeById(context.Background(), int64(employeeId))
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get employee by id", "Sukses mendapatkan karyawan berdasarkan id", respData)
}

func (e EmployeeHandler) GetEmployeeByName(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	employeeName := c.Query("employeeName")

	respData, errData := e.Usecase.Core.Employee.GetEmployeeByName(context.Background(), employeeName)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get employee by name", "Sukses mendapatkan karyawan berdasarkan nama", respData, 1, 0, 1)
}

func (e EmployeeHandler) Update(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	dataReq := new(models.UpdateEmployeeRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := e.Usecase.Core.Employee.Update(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success update employee", "Sukses mengubah karyawan", nil)
}

func (e EmployeeHandler) Delete(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	employeeId, err := strconv.Atoi(c.Query("employeeId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("employeeId"), cm.ErrValuePageIntInd("employeeId"), nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := e.Usecase.Core.Employee.Delete(context.Background(), int64(employeeId), username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success delete employee", "Sukses menghapus karyawan", nil)
}

func (e EmployeeHandler) Login(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	dataReq := new(models.LoginRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	respData, errData := e.Usecase.Core.Employee.Login(context.Background(), dataReq)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse(init, fiber.StatusOK, "Success login", "Sukses login", respData)
}

func (e EmployeeHandler) ForgotPassword(c *fiber.Ctx) error {
	init := exception.InitException(c, e.Conf, e.Log)

	dataReq := new(models.ForgotPasswordRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := e.Usecase.Core.Employee.ForgotPass(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse(init, fiber.StatusOK, "Success forgot password", "Sukses lupa password", nil)
}
