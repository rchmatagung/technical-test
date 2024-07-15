package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/location/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"fmt"
	"strings"
	"time"
	myvalidator "boilerplate/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Create(ctx context.Context, dataReq *models.CreateLocationRequest, createdBy string) (*models.GetLocation, *exception.Error)
	GetAllLocation(ctx context.Context) (*[]models.ListLocation, int, *exception.Error)
	GetLocationById(ctx context.Context, id int64) (*models.GetLocation, *exception.Error)
	GetLocationByName(ctx context.Context, LocationName string) (*models.GetLocation, *exception.Error)
	Update(ctx context.Context, dataReq *models.UpdateLocationRequest, updatedBy string) *exception.Error
	Delete(ctx context.Context, locationId int64, deletedBy string) *exception.Error
}

type LocationUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewLocationUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) LocationUsecase {
	return LocationUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (l LocationUsecase) Create(ctx context.Context, dataReq *models.CreateLocationRequest, createdBy string) (*models.GetLocation, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := l.Repo.Core.Location.GetLocationByName(ctx, dataReq.LocationName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data.LocationId != 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, fmt.Sprintf("Role with name %v already exist", dataReq.LocationName), fmt.Sprintf("Role dengan nama %v sudah ada", dataReq.LocationName))
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.LocationName, createdBy)

	id, err := l.Repo.Core.Location.Insert(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if id == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to insert data", "Gagal menyimpan data")
	}

	respData, err := l.Repo.Core.Location.GetLocationById(ctx, id)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if respData != nil {
		if respData.LocationId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get data", "Gagal mendapatkan data")
		}
	}

	return respData, nil
}

func (l LocationUsecase) GetAllLocation(ctx context.Context) (*[]models.ListLocation, int, *exception.Error) {
	data, count, err := l.Repo.Core.Location.GetAllLocation(ctx)
	if err != nil {
		return nil, 0, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if count == 0 {
		l.Log.Warn(`msg="Location not found" repo:"List Location"`)
		return nil, 0, exception.NewError(fiber.StatusOK, "Success get list location but 0 row", "Sukses mendapatkan list lokasi namun 0 row")
	}
	return data, count, nil
}

func (l LocationUsecase) GetLocationById(ctx context.Context, id int64) (*models.GetLocation, *exception.Error) {

	if id == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Id is required", "Id tidak boleh kosong")
	}

	data, err := l.Repo.Core.Location.GetLocationById(ctx, id)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.LocationId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Location not found", "Lokasi tidak ditemukan")
		}
	}
	return data, nil
}

func (l LocationUsecase) GetLocationByName(ctx context.Context, LocationName string) (*models.GetLocation, *exception.Error) {

	if LocationName == "" {
		return nil, exception.NewError(fiber.StatusBadRequest, "Name is required", "Nama tidak boleh kosong")
	}

	data, err := l.Repo.Core.Location.GetLocationByName(ctx, LocationName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.LocationId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get data", "Gagal mendapatkan data")
		}
	}
	return data, nil
}

func (l LocationUsecase) Update(ctx context.Context, dataReq *models.UpdateLocationRequest, updatedBy string) *exception.Error {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	eLocation, err := l.Repo.Core.Location.GetLocationById(ctx, dataReq.LocationId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if eLocation.LocationId == 0 {
		return exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Location with id %v not found", dataReq.LocationId), fmt.Sprintf("Lokasi dengan id %v tidak ditemukan", dataReq.LocationId))
	}

	dataLocation, err := l.Repo.Core.Location.GetLocationByName(ctx, dataReq.LocationName)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if dataLocation != nil {
		if strings.ToLower(dataLocation.LocationName) == strings.ToLower(eLocation.LocationName) {
			if strings.ToLower(dataReq.LocationName) == strings.ToLower(eLocation.LocationName) {
				return exception.NewError(fiber.StatusBadRequest, "Location is already exist", "Lokasi sudah ada")
			}
		}
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.LocationName, updatedBy, time.Now(), dataReq.LocationId)

	respData, err := l.Repo.Core.Location.Update(ctx, params...)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed update location with id %v", dataReq.LocationId), fmt.Sprintf("Gagal update lokasi dengan id %v", dataReq.LocationId))
	}

	return nil

}

func (l LocationUsecase) Delete(ctx context.Context, locationId int64, deletedBy string) *exception.Error {

	if locationId == 0 {
		return exception.NewError(fiber.StatusBadRequest, "Id is required", "Id tidak boleh kosong")
	}

	eLocation, err := l.Repo.Core.Location.GetLocationById(ctx, locationId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if eLocation.LocationId == 0 {
		return exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Location with id %v not found", locationId), fmt.Sprintf("Lokasi dengan id %v tidak ditemukan", locationId))
	}

	id, err := l.Repo.Core.Location.Delete(ctx, locationId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if id == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed delete location with id %v", id), fmt.Sprintf("Gagal hapus lokasi dengan id %v", id))
	}
	return nil
}
