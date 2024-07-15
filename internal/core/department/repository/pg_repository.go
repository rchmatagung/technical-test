package repository

import (
	"boilerplate/internal/core/department/models"
	"boilerplate/pkg/infra/db"
	"context"
)

type Repository interface {
	Insert(ctx context.Context, params ...interface{}) (int64, error)
	GetAllDepartment(ctx context.Context) (*[]models.ListDepartment, int, error)
	GetDepartmentById(ctx context.Context, id int64) (*models.GetDepartment, error)
	GetDepartmentByName(ctx context.Context, name string) (*models.GetDepartment, error)
	Update(ctx context.Context, params ...interface{}) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type DepartmentRepo struct {
	DBList *db.DatabaseList
}

func NewDepartmentRepo(dbList *db.DatabaseList) DepartmentRepo {
	return DepartmentRepo{
		DBList: dbList,
	}
}

func (d DepartmentRepo) Insert(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := d.DBList.DatabaseApp.Raw(InsertDepartment+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (d DepartmentRepo) GetAllDepartment(ctx context.Context) (*[]models.ListDepartment, int, error) {
	var response []models.ListDepartment
	var count int
	var err error
	var Where = "WHERE deleted_at IS NULL"
	err = d.DBList.DatabaseApp.Raw(qCount + Where).Scan(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = d.DBList.DatabaseApp.Raw(qSelectDepartment + Where).Scan(&response).Error
	return &response, count, err
}

func (d DepartmentRepo) GetDepartmentById(ctx context.Context, id int64) (*models.GetDepartment, error) {
	var response models.GetDepartment
	var Where = "WHERE department_id = ? AND deleted_at IS NULL"
	err := d.DBList.DatabaseApp.Raw(qSelectDepartment+Where, id).Scan(&response).Error
	return &response, err
}

func (d DepartmentRepo) GetDepartmentByName(ctx context.Context, name string) (*models.GetDepartment, error) {
	var response models.GetDepartment
	err := d.DBList.DatabaseApp.Raw(qSelectDepartment+qWhere+qDepartmentNameCaseInSensitive+qAnd+qNotDeleted, name).Scan(&response).Error
	return &response, err
}

func (d DepartmentRepo) Update(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := d.DBList.DatabaseApp.Raw(UpdateDepartment+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (d DepartmentRepo) Delete(ctx context.Context, id int64) (int64, error) {
	var response int64
	err := d.DBList.DatabaseApp.Raw(DeleteDepartment+qReturnID, id).Scan(&response).Error
	return response, err
}
