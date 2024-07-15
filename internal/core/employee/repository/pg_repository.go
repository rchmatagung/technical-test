package repository

import (
	"boilerplate/internal/core/employee/models"
	"boilerplate/pkg/infra/db"
	"context"
)

type Repository interface {
	Insert(ctx context.Context, params ...interface{}) (int64, error)
	GetAllEmployee(ctx context.Context) (*[]models.ListEmployeeScan, int, error)
	GetEmployeeById(ctx context.Context, id int64) (*models.GetEmployeeScan, error)
	GetEmployeeByName(ctx context.Context, name string) (*models.GetEmployeeScan, error)
	GetEmployeeByCode(ctx context.Context, code string) (*models.GetEmployeeScan, error)
	GetEmployeeByNameAndDeptIdAndPositionId(ctx context.Context, name string, deptId int, positionId int) (*models.GetEmployeeScan, error)
	Update(ctx context.Context, params ...interface{}) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	ForgotPass(ctx context.Context, params ...interface{}) (int64, error)
}

type EmployeeRepo struct {
	DBList *db.DatabaseList
}

func NewEmployeeRepo(dbList *db.DatabaseList) EmployeeRepo {
	return EmployeeRepo{
		DBList: dbList,
	}
}

func (e EmployeeRepo) Insert(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := e.DBList.DatabaseApp.Raw(qInsertEmployee+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (e EmployeeRepo) GetAllEmployee(ctx context.Context) (*[]models.ListEmployeeScan, int, error) {
	var response []models.ListEmployeeScan
	var count int
	var err error
	var Where = "Where e.deleted_at IS NULL"
	err = e.DBList.DatabaseApp.Raw(qCount+Where).Scan(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = e.DBList.DatabaseApp.Raw(qSelectEmployeeWithoutPassword+Where).Scan(&response).Error
	return &response, count, err
}

func (e EmployeeRepo) GetEmployeeById(ctx context.Context, id int64) (*models.GetEmployeeScan, error) {
	var response models.GetEmployeeScan
	var Where = "WHERE employee_id = ? AND e.deleted_at IS NULL"
	err := e.DBList.DatabaseApp.Raw(qSelectEmployeeWithoutPassword+Where, id).Scan(&response).Error
	return &response, err
}

func (e EmployeeRepo) GetEmployeeByName(ctx context.Context, name string) (*models.GetEmployeeScan, error) {
	var response models.GetEmployeeScan
	err := e.DBList.DatabaseApp.Raw(qSelectEmployeeWithoutPassword+qWhere+qEmployeeNameCaseInSensitive+qAnd+qNotDeleted, name).Scan(&response).Error
	return &response, err
}

func (e EmployeeRepo) GetEmployeeByCode(ctx context.Context, code string) (*models.GetEmployeeScan, error) {
	var response models.GetEmployeeScan
	var Where = "WHERE employee_code = ? AND e.deleted_at IS NULL"
	err := e.DBList.DatabaseApp.Raw(qSelectEmployee+Where, code).Scan(&response).Error
	return &response, err
}

func (e EmployeeRepo) GetEmployeeByNameAndDeptIdAndPositionId(ctx context.Context, name string, deptId int, positionId int) (*models.GetEmployeeScan, error) {
	var response models.GetEmployeeScan
	err := e.DBList.DatabaseApp.Raw(qSelectEmployeeWithoutPassword+qWhere+qEmployeeNameCaseInSensitive+qAnd+qDeptId+qAnd+qPositionId+qAnd+qNotDeleted, name, deptId, positionId).Scan(&response).Error
	return &response, err
}

func (e EmployeeRepo) Update(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := e.DBList.DatabaseApp.Raw(qUpdateEmployee+qReturnID, params...).Scan(&response).Error
	return response, err
}

func (e EmployeeRepo) Delete(ctx context.Context, id int64) (int64, error) {
	var response int64
	err := e.DBList.DatabaseApp.Raw(qDeleteEmployee+qReturnID, id).Scan(&response).Error
	return response, err
}

func (e EmployeeRepo) ForgotPass(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	err := e.DBList.DatabaseApp.Raw(qForgotPassword+qReturnID, params...).Scan(&response).Error
	return response, err
}