package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/employee/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"boilerplate/pkg/utils"
	"context"
	"fmt"
	myvalidator "boilerplate/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Create(ctx context.Context, dataReq *models.CreateEmployeeRequest, createdBy string) (*models.GetEmployeeScan, *exception.Error)
	GetAllEmployee(ctx context.Context) (*[]models.ListEmployeeScan, int, *exception.Error)
	GetEmployeeById(ctx context.Context, employeeId int64) (*models.GetEmployeeScan, *exception.Error)
	GetEmployeeByName(ctx context.Context, employeeName string) (*models.GetEmployeeScan, *exception.Error)
	Update(ctx context.Context, dataReq *models.UpdateEmployeeRequest, updatedBy string) *exception.Error
	Delete(ctx context.Context, employeeId int64, deletedBy string) *exception.Error
	Login(ctx context.Context, dataReq *models.LoginRequest) (*models.LoginResponse, *exception.Error)
	ForgotPass(ctx context.Context, dataReq *models.ForgotPasswordRequest, updatedBy string) *exception.Error
}

type EmployeeUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewEmployeeUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) EmployeeUsecase {
	return EmployeeUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (e EmployeeUsecase) Create(ctx context.Context, dataReq *models.CreateEmployeeRequest, createdBy string) (*models.GetEmployeeScan, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)

	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := e.Repo.Core.Employee.GetEmployeeByNameAndDeptIdAndPositionId(ctx, dataReq.EmployeeName, int(dataReq.DepartmentId), int(dataReq.PositionId))
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data.EmployeeId != 0 {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Employee with name %v already exist in department %v", dataReq.EmployeeName, data.DepartmentName), fmt.Sprintf("Employee dengan nama %v sudah ada pada departmen %v", dataReq.EmployeeName, data.DepartmentName))
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.EmployeeName, dataReq.Password, dataReq.DepartmentId, dataReq.PositionId, dataReq.Superior, createdBy)

	// Hashing password before insert to database dan replace password with hashed password
	hashedPassword, err := utils.HashingPassword(dataReq.Password)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	params[1] = hashedPassword

	respId, err := e.Repo.Core.Employee.Insert(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}
	if respId == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to create employee", "Gagal membuat employee")
	}

	respData, err := e.Repo.Core.Employee.GetEmployeeById(ctx, respId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData != nil {
		if respData.EmployeeId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
		}
	}

	return respData, nil
}

func (e EmployeeUsecase) GetAllEmployee(ctx context.Context) (*[]models.ListEmployeeScan, int, *exception.Error) {
	listEmployee, count, err := e.Repo.Core.Employee.GetAllEmployee(ctx)
	if err != nil {
		return nil, 0, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if count == 0 {
		e.Log.Warn(`msg="Employee Not Found" repo:"ListEmployee"`)
		return nil, 0, exception.NewError(fiber.StatusOK, "Success Get List Employee but 0 row", "Berhasil mendapatkan list employee tetapi 0 row")
	}

	return listEmployee, count, nil
}

func (e EmployeeUsecase) GetEmployeeById(ctx context.Context, employeeId int64) (*models.GetEmployeeScan, *exception.Error) {

	if employeeId == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Employee Id is required", "Employee Id harus diisi")
	}

	data, err := e.Repo.Core.Employee.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data != nil {
		if data.EmployeeId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
		}
	}

	return data, nil
}

func (e EmployeeUsecase) GetEmployeeByName(ctx context.Context, employeeName string) (*models.GetEmployeeScan, *exception.Error) {

	if employeeName == "" {
		return nil, exception.NewError(fiber.StatusBadRequest, "Employee Name is required", "Nama Pegawai harus diisi")
	}

	data, err := e.Repo.Core.Employee.GetEmployeeByName(ctx, employeeName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data != nil {
		if data.EmployeeId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
		}
	}

	return data, nil
}

func (e EmployeeUsecase) Update(ctx context.Context, dataReq *models.UpdateEmployeeRequest, updatedBy string) *exception.Error {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	eEmployee, err := e.Repo.Core.Employee.GetEmployeeById(ctx, dataReq.EmployeeId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if eEmployee.EmployeeId == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.EmployeeName, dataReq.DepartmentId, dataReq.PositionId, updatedBy, dataReq.EmployeeId)

	respData, err := e.Repo.Core.Employee.Update(ctx, params...)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to update employee", "Gagal mengupdate employee")
	}

	return nil
}

func (e EmployeeUsecase) Delete(ctx context.Context, employeeId int64, deletedBy string) *exception.Error {

	if employeeId == 0 {
		return exception.NewError(fiber.StatusBadRequest, "Employee Id is required", "Employee Id harus diisi")
	}

	data, err := e.Repo.Core.Employee.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data.EmployeeId == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	respData, err := e.Repo.Core.Employee.Delete(ctx, employeeId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to delete employee", "Gagal menghapus employee")
	}

	return nil
}

func (e EmployeeUsecase) Login(ctx context.Context, dataReq *models.LoginRequest) (*models.LoginResponse, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := e.Repo.Core.Employee.GetEmployeeByCode(ctx, dataReq.EmployeeCode)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data == nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	if data.EmployeeId == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	if dataReq.Password == "" {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Password required", "Password harus diisi")
	}

	if utils.CheckHashedPassword(data.Password, dataReq.Password) == false {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Password Not Match", "Password tidak cocok")
	}

	token, err := utils.GenereateJWT(e.Conf, dataReq.EmployeeCode, dataReq.RoleName)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	respData := &models.LoginResponse{
		EmployeeCode: dataReq.EmployeeCode,
		EmployeeName: data.EmployeeName,
		RoleName:     dataReq.RoleName,
		AccessToken:  token,
	}

	return respData, nil

}

func (e EmployeeUsecase) ForgotPass(ctx context.Context, dataReq *models.ForgotPasswordRequest, updatedBy string) *exception.Error {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)
	if errMsg != "" || errMsgInd != "" {
		return exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := e.Repo.Core.Employee.GetEmployeeByCode(ctx, dataReq.EmployeeCode)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data == nil {
		return exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	if data.EmployeeId == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Employee Not Found", "Pegawai tidak ditemukan")
	}

	if data.Password == "" {
		return exception.NewError(fiber.StatusInternalServerError, "Password required", "Password harus diisi")
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.Password, updatedBy, data.EmployeeCode)

	hashPassword, err := utils.HashingPassword(dataReq.Password)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	params[0] = hashPassword

	respData, err := e.Repo.Core.Employee.ForgotPass(ctx, params...)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to update password", "Gagal mengupdate password")
	}

	return nil

}
