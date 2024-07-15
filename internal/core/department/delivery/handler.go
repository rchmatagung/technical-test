package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/department/models"
	"boilerplate/internal/wrapper/usecase"
	cm "boilerplate/pkg/constants/message"
	"boilerplate/pkg/exception"
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

func (h DepartmentHandler) Create(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	dataReq := new(models.CreateDepartmentRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	data, errData := h.Usecase.Core.Department.Create(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create department", "Sukses membuat departemen", data)

}

func (h DepartmentHandler) GetAllDepartment(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	data, total, errData := h.Usecase.Core.Department.GetAllDepartment(context.Background())
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get all department", "Sukses mendapatkan semua departemen", data, 1, 0, total)
}

func (h DepartmentHandler) GetDepartmentById(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("id"), cm.ErrValuePageInt("id"), nil)
	}

	data, errData := h.Usecase.Core.Department.GetDepartmentById(context.Background(), int64(id))
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get department with id %v", id), fmt.Sprintf("Sukses mendapatkan departemen dengan id %v", id), data)
}

func (h DepartmentHandler) GetDepartmentByName(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	departmentName := c.Query("departmentName")

	data, errData := h.Usecase.Core.Department.GetDepartmentByName(context.Background(), departmentName)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get department with name %v", departmentName), fmt.Sprintf("Sukses mendapatkan departemen dengan nama %v", departmentName), data)
}

func (h DepartmentHandler) Update(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	dataReq := new(models.UpdateDepartmentRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := h.Usecase.Core.Department.Update(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}
	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success update department", "Sukses mengubah departemen", nil)
}

func (h DepartmentHandler) Delete(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	departmentId, err := strconv.Atoi(c.Query("departmentId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("departmentId"), cm.ErrValuePageInt("departmentId"), nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := h.Usecase.Core.Department.Delete(context.Background(), int64(departmentId), username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}
	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success delete department with id %v", departmentId), fmt.Sprintf("Sukses menghapus departemen dengan id %v", departmentId), nil)
}
