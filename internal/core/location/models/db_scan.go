package models

import "time"

type GetLocation struct {
	LocationId   int64      `json:"location_id"`
	LocationName string     `json:"location_name"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedBy    *string    `json:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type ListLocation struct {
	LocationId   int64      `json:"location_id"`
	LocationName string     `json:"location_name"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedBy    *string    `json:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
