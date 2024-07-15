package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/attendance/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"fmt"
	"time"

	myvalidator "boilerplate/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	AbsentIn(ctx context.Context, dataReq *models.CreateAttendanceAbsentInRequest, createdBy string) (*models.GetAttendanceScan, *exception.Error)
	AbsentOut(ctx context.Context, dataReq *models.CreateAttendanceAbsentOutRequest, updatedBy string) (*models.GetAttendanceScan, *exception.Error)
	GetAllAttendance(ctx context.Context) (*[]models.ListAttendanceScan, int, *exception.Error)
	GetAttendanceById(ctx context.Context, attendanceId int64) (*models.GetAttendanceScan, *exception.Error)
	GetAttendanceByEmployeeId(ctx context.Context, employeeId int64) (*models.GetAttendanceScan, *exception.Error)
	Delete(ctx context.Context, attendanceId int64, deletedBy string) *exception.Error
}

type AttendanceUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewAttendanceUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) AttendanceUsecase {
	return AttendanceUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (e AttendanceUsecase) AbsentIn(ctx context.Context, dataReq *models.CreateAttendanceAbsentInRequest, createdBy string) (*models.GetAttendanceScan, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)

	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := e.Repo.Core.Attendance.GetAttendanceByEmployeeId(ctx, dataReq.EmployeeId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data.AttendanceId != 0 {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Employee with id %v already absent in", dataReq.EmployeeId), fmt.Sprintf("Karyawan dengan id %v sudah absen masuk", dataReq.EmployeeId))
	}

	// data.absentin tidak boleh lebih dari data.absenout
	if data.AbsentOut.After(data.AbsentIn) {
		return nil, exception.NewError(fiber.StatusBadRequest, "Absent in cannot be greater than absent out", "Absen masuk tidak boleh lebih dari absen keluar")
	}

	params := make([]interface{}, 0)
	params = append(params, dataReq.EmployeeId, dataReq.LocationId, time.Now(), createdBy)

	attendanceId, err := e.Repo.Core.Attendance.InsertAbsentIn(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if attendanceId == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to create absent in", "Gagal membuat absen masuk")
	}

	respData, err := e.Repo.Core.Attendance.GetAttendanceById(ctx, attendanceId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData != nil {
		if respData.AttendanceId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get absent in", "Gagal mendapatkan absen masuk")
		}
	}

	return respData, nil
}

func (e AttendanceUsecase) AbsentOut(ctx context.Context, dataReq *models.CreateAttendanceAbsentOutRequest, updateBy string) (*models.GetAttendanceScan, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(dataReq)

	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	data, err := e.Repo.Core.Attendance.GetAttendanceByEmployeeId(ctx, dataReq.EmployeeId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data.AttendanceId == 0 {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Employee with id %v not found", dataReq.EmployeeId), fmt.Sprintf("Karyawan dengan id %v tidak ditemukan", dataReq.EmployeeId))
	}

	if data.AbsentIn.IsZero() {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Employee with id %v not absent in", dataReq.EmployeeId), fmt.Sprintf("Karyawan dengan id %v belum absen masuk", dataReq.EmployeeId))
	}

	if data.AbsentOut.After(data.AbsentIn) {
		return nil, exception.NewError(fiber.StatusNotFound, fmt.Sprintf("Employee with id %v already absent out", dataReq.EmployeeId), fmt.Sprintf("Karyawan dengan id %v sudah absen keluar", dataReq.EmployeeId))
	}

	params := make([]interface{}, 0)
	params = append(params, time.Now(), updateBy, dataReq.EmployeeId, dataReq.LocationId)

	attendanceId, err := e.Repo.Core.Attendance.InsertAbsentOut(ctx, params...)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if attendanceId == 0 {
		return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to create absent out", "Gagal membuat absen keluar")
	}

	respData, err := e.Repo.Core.Attendance.GetAttendanceById(ctx, attendanceId)

	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())

	}

	if respData != nil {
		if respData.AttendanceId == 0 {
			return nil, exception.NewError(fiber.StatusInternalServerError, "Failed to get absent out", "Gagal mendapatkan absen keluar")
		}
	}

	return respData, nil
}

func (e AttendanceUsecase) GetAllAttendance(ctx context.Context) (*[]models.ListAttendanceScan, int, *exception.Error) {

	listAttendance, count, err := e.Repo.Core.Attendance.GetAllAttendance(ctx)
	if err != nil {
		return nil, 0, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if count == 0 {
		return nil, 0, exception.NewError(fiber.StatusOK, "Success but no data found", "Sukses tetapi tidak ada data yang ditemukan")
	}

	return listAttendance, count, nil
}

func (e AttendanceUsecase) GetAttendanceById(ctx context.Context, attendanceId int64) (*models.GetAttendanceScan, *exception.Error) {

	if attendanceId == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Attendance id is required", "Id absen dibutuhkan")
	}

	data, err := e.Repo.Core.Attendance.GetAttendanceById(ctx, attendanceId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data != nil {
		if data.AttendanceId == 0 {
			return nil, exception.NewError(fiber.StatusNotFound, "Attendance not found", "Absen tidak ditemukan")
		}
	}

	return data, nil
}

func (e AttendanceUsecase) GetAttendanceByEmployeeId(ctx context.Context, employeeId int64) (*models.GetAttendanceScan, *exception.Error) {

	if employeeId == 0 {
		return nil, exception.NewError(fiber.StatusBadRequest, "Employee id is required", "Id karyawan dibutuhkan")
	}

	data, err := e.Repo.Core.Attendance.GetAttendanceByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data != nil {
		if data.AttendanceId == 0 {
			return nil, exception.NewError(fiber.StatusNotFound, "Attendance not found", "Absen tidak ditemukan")
		}
	}

	return data, nil
}

func (e AttendanceUsecase) Delete(ctx context.Context, attendanceId int64, deletedBy string) *exception.Error {

	if attendanceId == 0 {
		return exception.NewError(fiber.StatusBadRequest, "Attendance id is required", "Id absen dibutuhkan")
	}

	data, err := e.Repo.Core.Attendance.GetAttendanceById(ctx, attendanceId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if data != nil {
		if data.AttendanceId == 0 {
			return exception.NewError(fiber.StatusNotFound, "Attendance not found", "Absen tidak ditemukan")
		}
	}

	respData, err := e.Repo.Core.Attendance.Delete(ctx, attendanceId)
	if err != nil {
		return exception.NewError(fiber.StatusInternalServerError, err.Error(), err.Error())
	}

	if respData == 0 {
		return exception.NewError(fiber.StatusInternalServerError, "Failed to delete attendance", "Gagal menghapus absen")
	}
	return nil
}
