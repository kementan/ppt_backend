package role

import (
	"time"
)

type (
	RoleRequest struct {
		Name string `json:"name" binding:"required"`
	}

	RoleCreate struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RoleResponse struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
