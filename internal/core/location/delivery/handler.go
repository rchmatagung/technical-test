package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/location/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
	"context"
	"fmt"
	"strconv"

	cm "boilerplate/pkg/constants/message"

	"github.com/gofiber/fiber/v2"
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

func (h LocationHandler) Create(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	dataReq := new(models.CreateLocationRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	data, errData := h.Usecase.Core.Location.Create(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success create location", "Sukses membuat lokasi", data)
}

func (h LocationHandler) GetAllLocation(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	data, total, errData := h.Usecase.Core.Location.GetAllLocation(context.Background())
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log_Page(init, fiber.StatusOK, "Success get all location", "Sukses mendapatkan semua lokasi", data, 1, 0, total)
}

func (h LocationHandler) GetLocationById(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("id"), cm.ErrValuePageIntInd("id"), nil)
	}

	data, errData := h.Usecase.Core.Location.GetLocationById(context.Background(), int64(id))
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get location with id %v", id), fmt.Sprintf("Sukses mendapatkan lokasi dengan id %v", id), data)

}

func (h LocationHandler) GetLocationByName(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	locationName := c.Query("locationName")

	data, errData := h.Usecase.Core.Location.GetLocationByName(context.Background(), locationName)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success get location with name %v", locationName), fmt.Sprintf("Sukses mendapatkan lokasi dengan nama %v", locationName), data)

}

func (h LocationHandler) Update(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	dataReq := new(models.UpdateLocationRequest)
	if err := c.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrBodyParser, cm.ErrBodyParserInd, nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := h.Usecase.Core.Location.Update(context.Background(), dataReq, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success update location", "Sukses mengubah lokasi", nil)
}

func (h LocationHandler) Delete(c *fiber.Ctx) error {
	init := exception.InitException(c, h.Conf, h.Log)

	locationId, err := strconv.Atoi(c.Query("locationId"))
	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, cm.ErrValuePageInt("locationId"), cm.ErrValuePageIntInd("locationId"), nil)
	}

	username := fmt.Sprintf("%v", c.Locals("employee_name"))

	errData := h.Usecase.Core.Location.Delete(context.Background(), int64(locationId), username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, fmt.Sprintf("Success delete location with id %v", locationId), fmt.Sprintf("Sukses menghapus lokasi dengan id %v", locationId), nil)
}
