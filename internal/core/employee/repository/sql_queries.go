package repository

const (
	// Membuat pada atribut employee_code menjadi format (yymm + 0001) dan seterusnya

	qInsertEmployee = `
	INSERT INTO employee(employee_code, employee_name, password, department_id, position_id, superior, created_by)
	VALUES (
		CONCAT(
			TO_CHAR(CURRENT_DATE, 'YYMM'),
			LPAD(
				COALESCE(
					(
						SELECT CAST(SUBSTR(employee_code, 5, 4) AS INTEGER) + 1
						FROM employee
						WHERE employee_code LIKE CONCAT(TO_CHAR(CURRENT_DATE, 'YYMM'), '%')
						ORDER BY employee_code DESC
						LIMIT 1
					),
					1
				)::VARCHAR,
				4,
				'0'
			)
		),
		?, ?, ?, ?, ?, ?
	)
	`

	qReturnID = `
	RETURNING employee_id
	`

	qCount = `
	SELECT COUNT(employee_id) AS total FROM employee e
	`

	qSelectEmployee = `
	SELECT employee_id, employee_code, employee_name, password, d.department_id, d.department_name, p.position_id, p.position_name, superior, e.created_by, e.created_at, e.updated_by, e.updated_at
	FROM employee e
	INNER JOIN department d ON d.department_id = e.department_id
	INNER JOIN position p ON p.position_id = e.position_id
	`
	
	qSelectEmployeeWithoutPassword = `
	SELECT employee_id, employee_code, employee_name, d.department_id, d.department_name, p.position_id, p.position_name, superior, e.created_by, e.created_at, e.updated_by, e.updated_at
	FROM employee e
	INNER JOIN department d ON d.department_id = e.department_id
	INNER JOIN position p ON p.position_id = e.position_id
	`

	qWhere = `
	WHERE
	`
	qAnd = `
	AND
	`
	qNotDeleted = `
	e.deleted_at IS NULL
	`
	qEmployeeNameCaseInSensitive = `
	LOWER(employee_name) = LOWER(?)
	`
	qDeptId = `
	d.department_id = ?
	`
	qPositionId = `
	p.position_id = ?
	`
	qUpdateEmployee=`
	UPDATE employee
	SET employee_name = ?, department_id = ?, position_id = ?, updated_by = ?
	WHERE employee_id = ?
	`

	qDeleteEmployee = `
	UPDATE employee
	SET deleted_at = NOW()
	WHERE employee_id = ?
	`

	qForgotPassword = `
	UPDATE employee
	SET password = ?, updated_by = ?
	WHERE employee_code = ?
	`
)