package report_category

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_report_categories"

type (
	DataWithPagination struct {
		Row        []ReportCategoryResponse `json:"row"`
		Pagination util.PaginationResponse  `json:"pagination"`
	}

	ReportCategoryCreateRequest struct {
		Name string `json:"name" binding:"required"`
	}

	ReportCategoryUpdateRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	ReportCategoryCreate struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ReportCategoryUpdate struct {
		Name      sql.NullString `json:"name"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	ReportCategoryResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
