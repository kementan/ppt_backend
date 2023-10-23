package module_generator

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_land_statuses"

type (
	DataWithPagination struct {
		Row        []ModuleGeneratorResponse `json:"row"`
		Pagination util.PaginationResponse   `json:"pagination"`
	}

	ModuleGeneratorCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	ModuleGeneratorUpdateRequest struct {
		ID    string `json:"id" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	ModuleGeneratorCreate struct {
		Name      string    `json:"name"`
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ModuleGeneratorUpdate struct {
		Name      sql.NullString `json:"name"`
		Color     sql.NullString `json:"color"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	ModuleGeneratorResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
