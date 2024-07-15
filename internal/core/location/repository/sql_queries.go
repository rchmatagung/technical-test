package repository

const (
	InsertLocation = `INSERT INTO location(location_name, created_by) VALUES (?, ?)`
	
	qReturnID = `
	RETURNING location_id
	`

	qCount = `
	SELECT COUNT(location_id) AS total FROM location
	`

	qSelectLocation = `
	SELECT location_id, location_name, created_by, created_at, updated_by, updated_at
	FROM location
	`

	qWhere = `
	WHERE
	`

	qLocationNameCaseInSensitive = `
	LOWER(location_name) = LOWER(?)
	`

	qAnd = `
	AND
	`

	qNotDeleted = `
	deleted_at IS NULL
	`

	qUpdateLocation = `
	UPDATE location
	SET location_name = ?,
		updated_by = ?,
		updated_at = ?
	WHERE location_id = ?
	`
	
	qDeleteLocation = `
	UPDATE location
	SET deleted_at = NOW()
	WHERE location_id = ?
	`
)