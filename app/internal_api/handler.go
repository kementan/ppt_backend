package internal_api

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type InternalApiHandler struct {
	Usecase InternalApiUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase InternalApiUsecase, rdb *redis.Client) {
	handler := &InternalApiHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("fetch-all", handler.GetAll)

	// Fetching data and store into database
	// SIPDPS
	v1.GET("api-sipdps-laporan-tanam-fetch", handler.GetSIPDPSTanamFetch)
	v1.GET("api-sipdps-laporan-produktivitas-fetch", handler.GetSIPDPSProduktivitasFetch)
	v1.GET("api-sipdps-laporan-puso-fetch", handler.GetSIPDPSPusoFetch)
	v1.GET("api-sipdps-laporan-panen-fetch", handler.GetSIPDPSPanenFetch)
	// PERBENIHAN
	v1.GET("api-perbenihan-produsen-fetch", handler.GetPerbenihanProdusenFetch)
	v1.GET("api-perbenihan-rek-nas-fetch", handler.GetPerbenihanRekNasFetch)
	v1.GET("api-perbenihan-rek-bpsb-fetch", handler.GetPerbenihanRekBpsbFetch)
	v1.GET("api-perbenihan-rek-lssm-fetch", handler.GetPerbenihanRekLssmFetch)
	v1.GET("api-perbenihan-rek-penyaluran-fetch", handler.GetPerbenihanRekPenyaluranFetch)
	v1.GET("api-perbenihan-rek-penyebaran-fetch", handler.GetPerbenihanRekPenyebaranFetch)
	v1.GET("api-perbenihan-rek-produsen-fetch", handler.GetPerbenihanRekProdusenFetch)

	// Menampilkan data yang telah di fetch
	// SIPDPS
	v1.GET("api-sipdps-laporan-tanam", handler.GetSIPDPSTanam)
	v1.GET("api-sipdps-laporan-produktivitas", handler.GetSIPDPSProduktivitas)
	v1.GET("api-sipdps-laporan-puso", handler.GetSIPDPSPuso)
	v1.GET("api-sipdps-laporan-panen", handler.GetSIPDPSPanen)
	// PERBENIHAN
	v1.GET("api-perbenihan-produsen", handler.GetPerbenihanProdusen)
	v1.GET("api-perbenihan-rek-nas", handler.GetPerbenihanRekNas)
	v1.GET("api-perbenihan-rek-bpsb", handler.GetPerbenihanRekBpsb)
	v1.GET("api-perbenihan-rek-lssm", handler.GetPerbenihanRekLssm)
	v1.GET("api-perbenihan-rek-penyaluran", handler.GetPerbenihanRekPenyaluran)
	v1.GET("api-perbenihan-rek-penyebaran", handler.GetPerbenihanRekPenyebaran)

	// Langsung hit pada API

	// SIMLUH
	v1.GET("api-simluh-sertifikat", util.AuthMiddleware(handler.rdb), handler.GetSimluhSertifikat)
	v1.GET("api-simluh-riwayat-pelatihan", util.AuthMiddleware(handler.rdb), handler.GetSimluhRiwayatPelatihan)
}

func (handler *InternalApiHandler) GetAll(c *gin.Context) {
	handler.Usecase.GetAll(c)
}

// SIPDPS
func (handler *InternalApiHandler) GetSIPDPSTanamFetch(c *gin.Context) {
	handler.Usecase.GetSIPDPSTanamFetch(c)
}

func (handler *InternalApiHandler) GetSIPDPSProduktivitasFetch(c *gin.Context) {
	handler.Usecase.GetSIPDPSProduktivitasFetch(c)
}

func (handler *InternalApiHandler) GetSIPDPSPusoFetch(c *gin.Context) {
	handler.Usecase.GetSIPDPSPusoFetch(c)
}

func (handler *InternalApiHandler) GetSIPDPSPanenFetch(c *gin.Context) {
	handler.Usecase.GetSIPDPSPanenFetch(c)
}

func (handler *InternalApiHandler) GetSIPDPSTanam(c *gin.Context) {
	handler.Usecase.GetSIPDPSTanam(c)
}

func (handler *InternalApiHandler) GetSIPDPSProduktivitas(c *gin.Context) {
	handler.Usecase.GetSIPDPSProduktivitas(c)
}

func (handler *InternalApiHandler) GetSIPDPSPuso(c *gin.Context) {
	handler.Usecase.GetSIPDPSPuso(c)
}

func (handler *InternalApiHandler) GetSIPDPSPanen(c *gin.Context) {
	handler.Usecase.GetSIPDPSPanen(c)
}

// Perbenihan
func (handler *InternalApiHandler) GetPerbenihanProdusenFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanProdusenFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekNasFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekNasFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekBpsbFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekBpsbFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekLssmFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekLssmFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekPenyaluranFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekPenyaluranFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekPenyebaranFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekPenyebaranFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekProdusenFetch(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekProdusenFetch(c)
}

func (handler *InternalApiHandler) GetPerbenihanProdusen(c *gin.Context) {
	handler.Usecase.GetPerbenihanProdusen(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekNas(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekNas(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekBpsb(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekBpsb(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekLssm(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekLssm(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekPenyaluran(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekPenyaluran(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekPenyebaran(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekPenyebaran(c)
}

func (handler *InternalApiHandler) GetPerbenihanRekProdusen(c *gin.Context) {
	handler.Usecase.GetPerbenihanRekProdusen(c)
}

func (handler *InternalApiHandler) GetSimluhSertifikat(c *gin.Context) {
	handler.Usecase.GetSimluhSertifikat(c)
}
func (handler *InternalApiHandler) GetSimluhRiwayatPelatihan(c *gin.Context) {
	handler.Usecase.GetSimluhRiwayatPelatihan(c)
}
