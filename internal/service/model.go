package service

import (
	"database/sql"
	"time"
)

var table = "ppt_services"

type (
	ServiceCreateRequest struct {
		Name string `json:"name" binding:"required"`
	}

	ServiceUpdateRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	ServiceDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	ServiceCreate struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ServiceUpdate struct {
		Name      sql.NullString `json:"name"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	ServiceResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
