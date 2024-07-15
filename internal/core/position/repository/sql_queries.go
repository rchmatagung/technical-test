package repository

const (
	qInsertPosition = `
	INSERT INTO position(department_id, position_name, created_by)
	VALUES (?, ?, ?)
	`
	qReturnID = `
	RETURNING position_id
	`
	qCount = `
	SELECT COUNT(position_id) AS total FROM position p
	INNER JOIN department d ON d.department_id = p.department_id
	`
	qSelectPosition = `
	SELECT position_id, d.department_id, d.department_name, position_name, p.created_by, p.created_at, p.updated_by, p.updated_at
	FROM position p
	INNER JOIN department d ON d.department_id = p.department_id
	`
	qWhere = `
	WHERE
	`
	qAnd = `
	AND
	`
	qPositionId = `
	position_id = ?
	`
	qNotDeleted = `
	p.deleted_at IS NULL
	`
	qPositionNameCaseInSensitive = `
	LOWER(position_name) = LOWER(?)
	`
	qDeptId = `
	p.department_id = ?
	`
	qUpdatePosition = `
	UPDATE position
	SET 
		department_id = ?,
		position_name = ?,
		updated_by = ?,
		updated_at = ?
	WHERE position_id = ?
	`
	qDeletePosition = `
	UPDATE position
	SET deleted_at = NOW()
	WHERE position_id = ?
	`
)