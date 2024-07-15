package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/position/models"
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
	Create(ctx context.Context, dataReq *models.CreatePositionRequest, createdBy string) (*models.GetPosition, *exception.Error)
	GetAllPosition(ctx context.Context, departmentId int) (*[]models.ListPosition, int, *exception.Error)
	GetPositionById(ctx context.Context, PositionId int64) (*models.GetPosition, *exception.Error)
	GetPositionByName(ctx context.Context, PositionName string) (*models.GetPosition, *exception.Error)
	Update(ctx context.Context, dataReq *models.UpdatePositionRequest, updatedBy string) *exception.Error
	Delete(ctx context.Context, PositionId int64, deletedBy string) *exception.Error
}

type PositionUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewPositionUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) PositionUsecase {
	return PositionUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (p PositionUsecase) Create(ctx context.Context, dataReq *models.CreatePositionRequest, createdBy string) (*models.GetPosition, *exception.Error) {
	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := p.Repo.Core.Position.GetPositionByNameAndDeptId(ctx, dataReq.PositionName, int(dataReq.DepartmentId))
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), "")
	}
	if data.PositionId != 0 {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Role with name %v already exist in department %v", dataReq.PositionName, data.DepartmentName), fmt.Sprintf("Role dengan nama %v sudah ada pada departmen %v", dataReq.PositionName, data.DepartmentName))
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.DepartmentId, dataReq.PositionName, createdBy)

	respId, err := p.Repo.Core.Position.Insert(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), "")
	}
	if respId == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to insert data", "Gagal menyimpan data")
	}

	respData, err := p.Repo.Core.Position.GetPositionById(ctx, respId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if respData != nil {
		if respData.PositionId == 0 {
			return nil, exception.NewError(fiber.StatusNotFound, "Position not found", "posisi tidak ditemukan")
		}
	}

	return respData, nil
}

func (p PositionUsecase) GetAllPosition(ctx context.Context, departmentId int) (*[]models.ListPosition, int, *exception.Error) {
	listPosition, count, err := p.Repo.Core.Position.GetAllPosition(ctx, departmentId)
	if err != nil {
		return nil, 0, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if count == 0 {
		p.Log.Warn(`msg="Position not found" repo:"List Position"`)
		return nil, 0, exception.NewError(fiber.StatusOK, "Success get list Position but 0 row", "Sukses mendapatkan list posisi namun 0 row")
	}

	return listPosition, count, nil
}

func (p PositionUsecase) GetPositionById(ctx context.Context, PositionId int64) (*models.GetPosition, *exception.Error) {

	if PositionId == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Position id is required", "id posisi harus diisi")
	}

	data, err := p.Repo.Core.Position.GetPositionById(ctx, PositionId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.PositionId == 0 {
			return nil, exception.NewError(fiber.StatusNotFound, "Position not found", "posisi tidak ditemukan")
		}
	}

	return data, nil
}

func (p PositionUsecase) GetPositionByName(ctx context.Context, PositionName string) (*models.GetPosition, *exception.Error) {

	if PositionName == "" {
		return nil, exception.NewError(fiber.StatusBadRequest, "Position name is required", "nama posisi harus diisi")
	}

	data, err := p.Repo.Core.Position.GetPositionByName(ctx, PositionName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.PositionId == 0 {
			return nil, exception.NewError(fiber.StatusNotFound, "Position not found", "posisi tidak ditemukan")
		}
	}

	return data, nil
}

func (p PositionUsecase) Update(ctx context.Context, dataReq *models.UpdatePositionRequest, updatedBy string) *exception.Error {
	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	ePosition, err := p.Repo.Core.Position.GetPositionById(ctx, dataReq.PositionId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if ePosition.PositionId == 0 {
		return exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Position with id %v not found", dataReq.PositionId), fmt.Sprintf("posisi dengan id %v tidak ditemukan", dataReq.PositionId))
	}

	dataPosition, err := p.Repo.Core.Position.GetPositionByName(ctx, dataReq.PositionName)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if dataPosition != nil {
		if strings.ToLower(dataReq.PositionName) != strings.ToLower(ePosition.PositionName) {
			if strings.ToLower(dataReq.PositionName) == strings.ToLower(ePosition.PositionName) {
				return exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Position with name %v already exist", dataReq.PositionName), fmt.Sprintf("posisi dengan nama %v sudah ada", dataReq.PositionName))
			}
		}
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.DepartmentId, dataReq.PositionName, updatedBy, time.Now(), dataReq.PositionId)

	respData, err := p.Repo.Core.Position.Update(ctx, params...)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed update Position with id %v", dataReq.PositionId), fmt.Sprintf("Gagal update posisi dengan id %v", dataReq.PositionId))
	}

	return nil
}

func (u PositionUsecase) Delete(ctx context.Context, PositionId int64, deletedBy string) *exception.Error {
	if PositionId == 0 {
		return exception.NewError(fiber.StatusBadRequest, "Position id is required", "id posisi harus diisi")
	}

	ePosition, err := u.Repo.Core.Position.GetPositionById(ctx, PositionId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if ePosition.PositionId == 0 {
		return exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Position with id %v not found", PositionId), fmt.Sprintf("posisi dengan id %v tidak ditemukan", PositionId))
	}

	id, err := u.Repo.Core.Position.Delete(ctx, PositionId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if id == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed delete Position with id %v", PositionId), fmt.Sprintf("Gagal hapus posisi dengan id %v", PositionId))
	}

	return nil
}
