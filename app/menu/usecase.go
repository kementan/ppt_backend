package menu

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
)

type (
	MenuUsecase interface {
		Export(c *gin.Context)
		GetTable(c *gin.Context)
		GetByID(c *gin.Context)
		Create(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
	}

	usecase struct {
		repo MenuRepository
	}
)

func NewUsecase(repo MenuRepository) MenuUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Export(c *gin.Context) {
	data, err := uc.repo.Read(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Slug")
	xlsx.SetCellValue(sheet1Name, "C1", "Sort")
	xlsx.SetCellValue(sheet1Name, "D1", "Status")
	xlsx.SetCellValue(sheet1Name, "E1", "Created")
	xlsx.SetCellValue(sheet1Name, "F1", "Updated")

	err = xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	for i, each := range data {
		isActive := "Inactive"
		if each.IsActive {
			isActive = "Active"
		}

		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.Slug)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each.Sort)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), isActive)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each.CreatedAt)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), each.UpdatedAt)
	}

	err = xlsx.SaveAs("./data-menu.xlsx")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	buf := new(bytes.Buffer)
	err = xlsx.Write(buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write Excel file"})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=file1.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())
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
	var req MenuCreateRequest

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

	arg := MenuCreate{
		Name:     req.Name,
		Slug:     req.Slug,
		Sort:     req.Sort,
		IsActive: req.IsActive,
	}

	data, err := uc.repo.Create(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Update(c *gin.Context) {
	var req MenuUpdateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	arg := MenuUpdate{
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
		Slug: sql.NullString{
			String: req.Slug,
			Valid:  true,
		},
		Sort: sql.NullInt16{
			Int16: req.Sort,
			Valid: true,
		},
		IsActive: req.IsActive,
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
