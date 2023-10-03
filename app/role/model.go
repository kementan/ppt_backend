package role

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_roles"

type (
	DataWithPagination struct {
		Row        []RoleResponse          `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	RoleCreateRequest struct {
		Name string `json:"name" binding:"required"`
	}

	RoleUpdateRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	RoleCreate struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RoleUpdate struct {
		Name      sql.NullString `json:"name"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	RoleResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		IsCreate  bool      `json:"is_create"`
		IsRead    bool      `json:"is_read"`
		IsUpdate  bool      `json:"is_update"`
		IsDelete  bool      `json:"is_delete"`
		IsPublic  bool      `json:"is_public"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
