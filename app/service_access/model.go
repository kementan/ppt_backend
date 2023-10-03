package service_access

import (
	"database/sql"
	"time"
)

var table = "ppt_service_accesss"

type (
	ServiceAccessCreateRequest struct {
		ServiceID string `json:"service_id" binding:"required"`
		RoleID    string `json:"role_id" binding:"required"`
	}

	ServiceAccessUpdateRequest struct {
		ID        string `json:"id" binding:"required"`
		ServiceID string `json:"service_id" binding:"required"`
		RoleID    string `json:"role_id" binding:"required"`
	}

	ServiceAccessDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	ServiceAccessCreate struct {
		ServiceID string    `json:"service_id"`
		RoleID    string    `json:"role_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ServiceAccessUpdate struct {
		ServiceID sql.NullString `json:"service_id"`
		RoleID    sql.NullString `json:"role_id"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	ServiceAccessResponse struct {
		HashedID        string    `json:"id"`
		HashedServiceID string    `json:"service_id"`
		HashedRoleID    string    `json:"role_id"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}
)
