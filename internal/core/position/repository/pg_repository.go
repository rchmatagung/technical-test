package repository

import (
	"boilerplate/internal/core/position/models"
	"boilerplate/pkg/infra/db"
	"context"
)

type Repository interface {
	Insert(ctx context.Context, params ...interface{}) (int64, error)
	GetAllPosition(ctx context.Context, departmentId int) (*[]models.ListPosition, int, error)
	GetPositionById(ctx context.Context, id int64) (*models.GetPosition, error)
	GetPositionByName(ctx context.Context, name string) (*models.GetPosition, error)
	Update(ctx context.Context, params ...interface{}) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	GetPositionByNameAndDeptId(ctx context.Context, name string, deptId int) (*models.GetPosition, error)
}

type PositionRepo struct {
	DBList *db.DatabaseList
}

func NewPositionRepo(dbList *db.DatabaseList) PositionRepo {
	return PositionRepo{
		DBList: dbList,
	}
}

func (p PositionRepo) Insert(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := p.DBList.DatabaseApp.Raw(qInsertPosition+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (p PositionRepo) GetAllPosition(ctx context.Context, departmentId int) (*[]models.ListPosition,int, error) {
	var response []models.ListPosition
	var count int
	var Where = "Where d.department_id = ? AND p.deleted_at IS NULL"
	var err error
	err = p.DBList.DatabaseApp.Raw(qCount+Where, departmentId).Scan(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = p.DBList.DatabaseApp.Raw(qSelectPosition+Where, departmentId).Scan(&response).Error
	return &response, count, err
}

func (p PositionRepo) GetPositionById(ctx context.Context, id int64) (*models.GetPosition, error) {
	var response models.GetPosition
	err := p.DBList.DatabaseApp.Raw(qSelectPosition+qWhere+qPositionId+qAnd+qNotDeleted, id).Scan(&response).Error
	return &response, err
}

func (p PositionRepo) GetPositionByName(ctx context.Context, name string) (*models.GetPosition, error) {
	var response models.GetPosition
	err := p.DBList.DatabaseApp.Raw(qSelectPosition+qWhere+qPositionNameCaseInSensitive+qAnd+qNotDeleted, name).Scan(&response).Error
	return &response, err
}

func (p PositionRepo) GetPositionByNameAndDeptId(ctx context.Context, name string, deptId int) (*models.GetPosition, error) {
	var response models.GetPosition
	err := p.DBList.DatabaseApp.Raw(qSelectPosition+qWhere+qPositionNameCaseInSensitive+qAnd+qDeptId+qAnd+qNotDeleted, name, deptId).Scan(&response).Error
	return &response, err
}

func (p PositionRepo) Update(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := p.DBList.DatabaseApp.Raw(qUpdatePosition+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (p PositionRepo) Delete(ctx context.Context, id int64) (int64, error) {
	var response int64
	err := p.DBList.DatabaseApp.Raw(qDeletePosition+qReturnID, id).Scan(&response).Error
	return response, err
}
