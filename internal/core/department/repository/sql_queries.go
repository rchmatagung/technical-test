package repository

const (
	InsertDepartment = `INSERT INTO department(department_name, created_by) VALUES (?, ?)`

	qReturnID = `
	RETURNING department_id
	`

	qCount = `
	SELECT COUNT(department_id) AS total FROM department
	`

	qSelectDepartment = `
	SELECT department_id, department_name, created_by, created_at, updated_by, updated_at
	FROM department
	`

	qWhere = `
	WHERE
	`

	qDepartmentNameCaseInSensitive = `
	LOWER(department_name) = LOWER(?)
	`

	qAnd = `
	AND
	`

	qNotDeleted = `
	deleted_at IS NULL
	`

	UpdateDepartment = `
	UPDATE department
	SET department_name = ?,
		updated_by = ?,
		updated_at = ?
	WHERE department_id = ?
	`

	DeleteDepartment = `
	UPDATE department
	SET deleted_at = NOW()
	WHERE department_id = ?
	`
)