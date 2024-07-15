package repository

const (
	qInsertAbsentIn = `INSERT INTO attendance (employee_id, location_id, absent_in, created_by) VALUES (?, ?, ?, ?)`

	qReturnID = `RETURNING attendance_id`

	// insert absent out query dengan menggunalan employee_id untuk where nya dan update absent_out
	qInsertAbsentOut=`
	UPDATE attendance SET
	absent_out = ?, updated_by = ?
	WHERE employee_id = ? AND location_id = ? AND deleted_at IS NULL
	`

	qCount = `SELECT COUNT(attendance_id) AS total FROM attendance a `

	qSelectAttendance = `
	SELECT attendance_id, e.employee_id, e.employee_name, l.location_id, l.location_name, absent_in, absent_out, a.created_by, a.created_at, a.updated_by, a.updated_at
	FROM attendance a
	INNER JOIN employee e ON e.employee_id = a.employee_id
	INNER JOIN location l ON l.location_id = a.location_id
	`

	qUpdateAttendance = `
	UPDATE attendance SET
	employee_id = ?, location_id = ?, absent_in = ?, absent_out = ?, updated_by = ?
	WHERE attendance_id = ?
	`

	qDeleteAttendance = `
	UPDATE attendance SET
	deleted_at = NOW()
	WHERE attendance_id = ?
	`
)