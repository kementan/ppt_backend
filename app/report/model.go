package report

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_reports"

type (
	DataWithPagination struct {
		Row        []LandStatusResponse    `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	LandStatusCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	LandStatusUpdateRequest struct {
		ID    string `json:"id" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	LandStatusCreate struct {
		Name      string    `json:"name"`
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	LandStatusUpdate struct {
		Name      sql.NullString `json:"name"`
		Color     sql.NullString `json:"color"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	LandStatusResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
