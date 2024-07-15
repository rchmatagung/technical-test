package models

type CreateLocationRequest struct {
	LocationName string `json:"location_name" validate:"required"`
}

type UpdateLocationRequest struct {
	LocationId   int64  `json:"location_id" validate:"required"`
	LocationName string `json:"location_name" validate:"required"`
}