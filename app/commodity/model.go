package commodity

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_commodities"

type (
	DataWithPagination struct {
		Row        []CommodityResponse     `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	CommodityCreateRequest struct {
		Name        string `json:"name" binding:"max=255"`
		Description string `json:"description"`
	}

	CommodityUpdateRequest struct {
		ID          string `json:"id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"max=255"`
	}

	CommodityCreate struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	CommodityUpdate struct {
		Name        sql.NullString `json:"name"`
		Description sql.NullString `json:"description"`
		CreatedAt   sql.NullString `json:"created_at"`
		UpdatedAt   sql.NullString `json:"updated_at"`
	}

	CommodityResponse struct {
		HashedID    string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
