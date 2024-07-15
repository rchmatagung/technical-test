package repository

import (
	"boilerplate/internal/core/attendance/models"
	"boilerplate/pkg/infra/db"
	"context"
)

type Repository interface {
	InsertAbsentIn(ctx context.Context, params ...interface{}) (int64, error)
	InsertAbsentOut(ctx context.Context, params ...interface{}) (int64, error)
	GetAllAttendance(ctx context.Context) (*[]models.ListAttendanceScan, int, error)
	GetAttendanceById(ctx context.Context, id int64) (*models.GetAttendanceScan, error)
	GetAttendanceByEmployeeId(ctx context.Context, employeeId int64) (*models.GetAttendanceScan, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type AttendanceRepo struct {
	DBList *db.DatabaseList
}

func NewAttendanceRepo(dbList *db.DatabaseList) AttendanceRepo {
	return AttendanceRepo{
		DBList: dbList,
	}
}

func (a AttendanceRepo) InsertAbsentIn(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := a.DBList.DatabaseApp.Raw(qInsertAbsentIn+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (a AttendanceRepo) InsertAbsentOut(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := a.DBList.DatabaseApp.Raw(qInsertAbsentOut+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (a AttendanceRepo) GetAllAttendance(ctx context.Context) (*[]models.ListAttendanceScan, int, error) {
	var response []models.ListAttendanceScan
	var count int
	var err error
	var Where = "Where a.deleted_at IS NULL"
	err = a.DBList.DatabaseApp.Raw(qCount+Where).Scan(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = a.DBList.DatabaseApp.Raw(qSelectAttendance+Where).Scan(&response).Error
	return &response, count, err
}

func (a AttendanceRepo) GetAttendanceById(ctx context.Context, id int64) (*models.GetAttendanceScan, error) {
	var response models.GetAttendanceScan
	var Where = "WHERE attendance_id = ? AND a.deleted_at IS NULL"
	err := a.DBList.DatabaseApp.Raw(qSelectAttendance+Where, id).Scan(&response).Error
	return &response, err
}

func (a AttendanceRepo) GetAttendanceByEmployeeId(ctx context.Context, employeeId int64) (*models.GetAttendanceScan, error) {
	var response models.GetAttendanceScan
	var Where = "WHERE e.employee_id = ? AND a.deleted_at IS NULL"
	err := a.DBList.DatabaseApp.Raw(qSelectAttendance+Where, employeeId).Scan(&response).Error
	return &response, err
}

func (a AttendanceRepo) Delete(ctx context.Context, id int64) (int64, error) {
	var response int64
	err := a.DBList.DatabaseApp.Raw(qDeleteAttendance+qReturnID, id).Scan(&response).Error
	return response, err
}

