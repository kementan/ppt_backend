package configuration

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_configurations"

type (
	DataWithPagination struct {
		Row        []ConfigurationResponse `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	ConfigurationCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Value string `json:"value" binding:"required"`
	}

	ConfigurationUpdateRequest struct {
		ID    string `json:"id" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Value string `json:"value" binding:"required"`
	}

	ConfigurationCreate struct {
		Name      string    `json:"name"`
		Value     string    `json:"value"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ConfigurationUpdate struct {
		Name      sql.NullString `json:"name"`
		Value     sql.NullString `json:"value"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	ConfigurationResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		Value     string    `json:"value"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
