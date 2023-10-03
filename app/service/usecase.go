package service

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
)

type (
	ServiceUsecase interface {
		Create(c *gin.Context)
		Read(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
	}

	usecase struct {
		repo ServiceRepository
	}
)

func NewUsecase(repo ServiceRepository) ServiceUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Create(c *gin.Context) {
	var req ServiceCreateRequest

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

	arg := ServiceCreate{
		Name:     req.Name,
		Image:    req.Image,
		IsActive: req.IsActive,
	}

	data, err := uc.repo.Create(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Read(c *gin.Context) {
	data, err := uc.repo.Read(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Update(c *gin.Context) {
	var req ServiceUpdateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	arg := ServiceUpdate{
		Name: sql.NullString{
			String: req.Name,
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
	var req ServiceDeleteRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	err = uc.repo.Delete(c, req.ID)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "data successfully deleted")
}
