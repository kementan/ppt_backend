package role

import (
	"database/sql"
	"time"
)

var table = "ppt_roles"

type (
	RoleRequest struct {
		Name string `json:"name" binding:"required"`
	}

	RoleUpdateRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	RoleDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	RoleCreate struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RoleDelete struct {
		ID int64 `json:"id" binding:"required,min=1"`
	}

	RoleUpdate struct {
		Name      sql.NullString `json:"name"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	RoleResponse struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
