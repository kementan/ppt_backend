package menu

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_menus"
var subtable = "ppt_sub_menus"

type (
	DataWithPagination struct {
		Row        []MenuResponse          `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	MenuCreateRequest struct {
		Name     string `json:"name" binding:"required"`
		Slug     string `json:"slug" binding:"required"`
		Sort     int    `json:"sort" binding:"required"`
		IsActive *bool  `json:"is_active"`
	}

	MenuUpdateRequest struct {
		ID       string `json:"id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Slug     string `json:"slug" binding:"required"`
		Sort     int16  `json:"sort" binding:"required"`
		IsActive *bool  `json:"is_active"`
	}

	MenuCreate struct {
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
		Sort      int       `json:"sort"`
		IsActive  *bool     `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	MenuUpdate struct {
		Name      sql.NullString `json:"name"`
		Slug      sql.NullString `json:"slug"`
		Sort      sql.NullInt16  `json:"sort"`
		IsActive  *bool          `json:"is_active"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	MenuResponse struct {
		HashedID  string    `json:"id"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
		Sort      int       `json:"sort"`
		IsActive  bool      `json:"is_active"`
		SubMenu   int       `json:"sub_menu"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
