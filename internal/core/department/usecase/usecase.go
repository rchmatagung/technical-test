package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/department/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	myvalidator "boilerplate/pkg/validator"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Create(ctx context.Context, dataReq *models.CreateDepartmentRequest, createdBy string) (*models.GetDepartment, *exception.Error)
	GetAllDepartment(ctx context.Context) (*[]models.ListDepartment, int, *exception.Error)
	GetDepartmentById(ctx context.Context, id int64) (*models.GetDepartment, *exception.Error)
	GetDepartmentByName(ctx context.Context, DepartmentName string) (*models.GetDepartment, *exception.Error)
	Update(ctx context.Context, dataReq *models.UpdateDepartmentRequest, updatedBy string) *exception.Error
	Delete(ctx context.Context, departmentId int64, deletedBy string) *exception.Error
}

type DepartmentUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewDepartmentUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) DepartmentUsecase {
	return DepartmentUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (d DepartmentUsecase) Create(ctx context.Context, dataReq *models.CreateDepartmentRequest, createdBy string) (*models.GetDepartment, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := d.Repo.Core.Department.GetDepartmentByName(ctx, dataReq.DepartmentName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data.DepartmentId != 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, fmt.Sprintf("Role with name %v already exist", dataReq.DepartmentName), fmt.Sprintf("Role dengan nama %v sudah ada", dataReq.DepartmentName))
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.DepartmentName, createdBy)

	id, err := d.Repo.Core.Department.Insert(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if id == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to insert data", "Gagal menyimpan data")
	}

	respData, err := d.Repo.Core.Department.GetDepartmentById(ctx, id)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData != nil {
		if respData.DepartmentId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get data", "Gagal mendapatkan data")
		}
	}

	return respData, nil

}

func (d DepartmentUsecase) GetAllDepartment(ctx context.Context) (*[]models.ListDepartment, int, *exception.Error) {
	data, count, err := d.Repo.Core.Department.GetAllDepartment(ctx)
	if err != nil {
		return nil, 0, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if count == 0 {
		d.Log.Warn(`msg="Department not found" repo:"List Department"`)
		return nil, 0, exception.NewError(fiber.StatusOK, "Success get list department but 0 row", "Sukses mendapatkan list departemen namun 0 row")
	}

	return data, count, nil
}

func (d DepartmentUsecase) GetDepartmentById(ctx context.Context, id int64) (*models.GetDepartment, *exception.Error) {

	if id == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Department Id is required", "Department Id tidak boleh kosong")
	}

	data, err := d.Repo.Core.Department.GetDepartmentById(ctx, id)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.DepartmentId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Department not found", "Departemen tidak ditemukan")
		}
	}

	return data, nil
}

func (d DepartmentUsecase) GetDepartmentByName(ctx context.Context, DepartmentName string) (*models.GetDepartment, *exception.Error) {

	if DepartmentName == "" {
		return nil, exception.NewError(fiber.StatusBadRequest, "Department Name is required", "Nama Departemen tidak boleh kosong")
	}

	data, err := d.Repo.Core.Department.GetDepartmentByName(ctx, DepartmentName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if data != nil {
		if data.DepartmentId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Department not found", "Departemen tidak ditemukan")
		}
	}

	return data, nil
}

func (d DepartmentUsecase) Update(ctx context.Context, dataReq *models.UpdateDepartmentRequest, updatedBy string) *exception.Error {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	eDepartment, err := d.Repo.Core.Department.GetDepartmentById(ctx, dataReq.DepartmentId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if eDepartment.DepartmentId == 0 {
		return exception.NewError(fiber.StatusBadRequest, fmt.Sprintf("Location with id %v not found", dataReq.DepartmentId), fmt.Sprintf("Lokasi dengan id %v tidak ditemukan", dataReq.DepartmentId))
	}

	dataDepartment, err := d.Repo.Core.Department.GetDepartmentByName(ctx, dataReq.DepartmentName)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if dataDepartment != nil {
		if strings.ToLower(dataDepartment.DepartmentName) == strings.ToLower(eDepartment.DepartmentName) {
			if strings.ToLower(dataReq.DepartmentName) == strings.ToLower(eDepartment.DepartmentName) {
				return exception.NewError(fiber.StatusBadRequest, "Department is already exist", "Departemen sudah ada")
			}
		}
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.DepartmentName, updatedBy, time.Now(), dataReq.DepartmentId)

	respData, err := d.Repo.Core.Department.Update(ctx, params...)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed update department with id %v", dataReq.DepartmentId), fmt.Sprintf("Gagal update departemen dengan id %v", dataReq.DepartmentId))
	}

	return nil
}

func (d DepartmentUsecase) Delete(ctx context.Context, departmentId int64, deletedBy string) *exception.Error {

	if departmentId == 0 {
		return exception.NewError(fiber.StatusBadRequest, "Department Id is required", "Department Id tidak boleh kosong")
	}

	eDepartment, err := d.Repo.Core.Department.GetDepartmentById(ctx, departmentId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if eDepartment.DepartmentId == 0 {
		return exception.NewError(fiber.StatusBadRequest, fmt.Sprintf("Department with id %v not found", departmentId), fmt.Sprintf("Departemen dengan id %v tidak ditemukan", departmentId))
	}

	id, err := d.Repo.Core.Department.Delete(ctx, departmentId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if id == 0 {
		return exception.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed delete department with id %v", departmentId), fmt.Sprintf("Gagal hapus departemen dengan id %v", departmentId))
	}

	return nil
}
