package repository

import (
	"boilerplate/internal/core/location/models"
	"boilerplate/pkg/infra/db"
	"context"
)

type Repository interface {
	Insert(ctx context.Context, params ...interface{}) (int64, error)
	GetAllLocation(ctx context.Context) (*[]models.ListLocation, int, error)
	GetLocationById(ctx context.Context, id int64) (*models.GetLocation, error)
	GetLocationByName(ctx context.Context, name string) (*models.GetLocation, error)
	Update(ctx context.Context, params ...interface{}) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type LocationRepo struct {
	DBList *db.DatabaseList
}

func NewLocationRepo(dbList *db.DatabaseList) LocationRepo {
	return LocationRepo{
		DBList: dbList,
	}
}

func (l LocationRepo) Insert(ctx context.Context, params ...interface{})(int64, error) {
	var response int64
	err := l.DBList.DatabaseApp.Raw(InsertLocation+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (l LocationRepo) GetAllLocation(ctx context.Context)(*[]models.ListLocation, int, error) {
	var response []models.ListLocation
	var count int
	var err error
	var Where = "WHERE deleted_at IS NULL"
	err = l.DBList.DatabaseApp.Raw(qCount+Where).Scan(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = l.DBList.DatabaseApp.Raw(qSelectLocation+Where).Scan(&response).Error
	return &response, count, err
}

func (l LocationRepo) GetLocationById(ctx context.Context, id int64)(*models.GetLocation, error) {
	var response models.GetLocation
	var Where = "WHERE location_id = ? AND deleted_at IS NULL"
	err := l.DBList.DatabaseApp.Raw(qSelectLocation+Where, id).Scan(&response).Error
	return &response, err
}

func (l LocationRepo) GetLocationByName(ctx context.Context, name string)(*models.GetLocation, error) {
	var response models.GetLocation
	err := l.DBList.DatabaseApp.Raw(qSelectLocation+qWhere+qLocationNameCaseInSensitive+qAnd+qNotDeleted, name).Scan(&response).Error
	return &response, err
}

func (l LocationRepo) Update(ctx context.Context, params ...interface{})(int64, error) {
	var response int64
	err := l.DBList.DatabaseApp.Raw(qUpdateLocation+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (l LocationRepo) Delete(ctx context.Context, id int64)(int64, error) {
	var response int64
	err := l.DBList.DatabaseApp.Raw(qDeleteLocation+qReturnID, id).Scan(&response).Error
	return response, err
}
