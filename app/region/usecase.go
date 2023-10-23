package region

import (
	"net/http"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
)

type (
	RegionUsecase interface {
		GetList(c *gin.Context)
		GetRegion(c *gin.Context)
	}

	usecase struct {
		repo RegionRepository
	}
)

func NewUsecase(repo RegionRepository) RegionUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetList(c *gin.Context) {
	level := c.Param("level")
	parentCode := c.Param("parentCode")

	data, err := uc.repo.GetList(c, level, parentCode)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetRegion(c *gin.Context) {
	code := c.Param("code")

	data, err := uc.repo.GetRegion(c, code)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}
