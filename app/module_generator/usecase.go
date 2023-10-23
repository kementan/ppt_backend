package module_generator

import (
	"database/sql"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
)

type (
	ModuleGeneratorUsecase interface {
		GetTable(c *gin.Context)
		GetList(c *gin.Context)
		GetByID(c *gin.Context)
		Create(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
	}

	usecase struct {
		repo ModuleGeneratorRepository
	}
)

func NewUsecase(repo ModuleGeneratorRepository) ModuleGeneratorUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetTable(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "desc")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	arg := util.DataFilter{
		Search:    search,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Page:      pageInt,
		PageSize:  pageSizeInt,
	}

	totalRecords, err := uc.repo.CountRecords(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSizeInt)))

	data, err := uc.repo.GetTable(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	pagination := util.PaginationResponse{
		CurrentPage:  pageInt,
		PageSize:     pageSizeInt,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	response := DataWithPagination{
		Row:        data,
		Pagination: pagination,
	}

	util.JOK(c, http.StatusOK, response)
}

func (uc *usecase) GetList(c *gin.Context) {
	data, err := uc.repo.Read(c)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetByID(c *gin.Context) {
	type ID struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	var req ID

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	id, _ := util.Decrypt(req.ID, "f")

	data, err := uc.repo.GetDataBy(c, "id", id)
	if err != nil {
		if err.Error() == "record not found" {
			util.JERR(c, http.StatusNotFound, errors.New("record not found"))
		} else {
			util.JERR(c, http.StatusInternalServerError, err)
		}
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Create(c *gin.Context) {
	var req ModuleGeneratorCreateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	exist, _ := uc.repo.GetDataBy(c, "name", req.Name)
	if exist.Name != "" {
		util.JERR(c, http.StatusConflict, errors.New("data already exists"))
		return
	}

	arg := ModuleGeneratorCreate{
		Name:  req.Name,
		Color: req.Color,
	}

	data, err := uc.repo.Create(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Update(c *gin.Context) {
	var req ModuleGeneratorUpdateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	arg := ModuleGeneratorUpdate{
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
		Color: sql.NullString{
			String: req.Color,
			Valid:  true,
		},
	}

	data, err := uc.repo.Update(c, req.ID, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Delete(c *gin.Context) {
	var req util.TableIDs

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	successIDs, failedIDs, err := uc.repo.Delete(c, req.IDs)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response := struct {
		SuccessIDs []string `json:"success_ids"`
		FailedIDs  []string `json:"failed_ids"`
	}{
		SuccessIDs: successIDs,
		FailedIDs:  failedIDs,
	}

	util.JOK(c, http.StatusOK, response)
}
