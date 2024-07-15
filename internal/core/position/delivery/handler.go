package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/position/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
	"context"
	"fmt"
	"strconv"

	cm "boilerplate/pkg/constants/message"

	"github.com/gofiber/fiber/v2"
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

func (p PositionHandler) Create(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	dataReq := new(models.CreatePositionRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	respData, errData := p.Usecase.Core.Position.Create(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create position", "Sukses membuat posisi", respData)
}

func (p PositionHandler) GetAllPosition(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	departmentId, err := strconv.Atoi(c.Query("department_id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("department_id"), cm.ErrValuePageIntInd("department_id"), nil)
	}

	respData, total, errData := p.Usecase.Core.Position.GetAllPosition(context.Background(), departmentId)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get all position", "Sukses mendapatkan semua posisi", respData, 1, 0, total)
}

func (p PositionHandler) GetPositionById(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	positionId, err := strconv.Atoi(c.Query("positionId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("positionId"), cm.ErrValuePageIntInd("positionId"), nil)
	}

	respData, errData := p.Usecase.Core.Position.GetPositionById(context.Background(), int64(positionId))
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get Position with id %v", positionId), fmt.Sprintf("Sukses mendapatkan posisi dengan id %v", positionId), respData)
}

func (p PositionHandler) GetPositionByName(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	positionName := c.Query("positionName")

	respData, errData := p.Usecase.Core.Position.GetPositionByName(context.Background(), positionName)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get Position with name %v", positionName), fmt.Sprintf("Sukses mendapatkan posisi dengan nama %v", positionName), respData)
}

func (p PositionHandler) Update(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	dataReq := new(models.UpdatePositionRequest)
	err := c.BodyParser(dataReq)
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := p.Usecase.Core.Position.Update(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success update position", "Sukses mengubah posisi", nil)
}

func (p PositionHandler) Delete(c *fiber.Ctx) error {
	init := exception.InitException(c, p.Conf, p.Log)

	positionId, err := strconv.Atoi(c.Query("positionId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("positionId"), cm.ErrValuePageIntInd("positionId"), nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := p.Usecase.Core.Position.Delete(context.Background(), int64(positionId), username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}
	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success delete position with id %v", positionId), fmt.Sprintf("Sukses menghapus posisi dengan id %v", positionId), nil)
}
