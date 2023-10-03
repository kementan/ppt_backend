package service_access

import (
	"database/sql"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
)

type (
	ServiceAccessUsecase interface {
		Create(c *gin.Context)
		Read(c *gin.Context)
		Update(c *gin.Context)
		Delete(c *gin.Context)
	}

	usecase struct {
		repo ServiceAccessRepository
	}
)

func NewUsecase(repo ServiceAccessRepository) ServiceAccessUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Create(c *gin.Context) {
	var req ServiceAccessCreateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	serviceID, _ := util.Decrypt(req.ServiceID, "f")
	roleID, _ := util.Decrypt(req.RoleID, "f")

	arg := ServiceAccessCreate{
		ServiceID: serviceID,
		RoleID:    roleID,
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
	var req ServiceAccessUpdateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	arg := ServiceAccessUpdate{
		ServiceID: sql.NullString{
			String: req.ServiceID,
			Valid:  true,
		},
		RoleID: sql.NullString{
			String: req.RoleID,
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
	var req ServiceAccessDeleteRequest

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
