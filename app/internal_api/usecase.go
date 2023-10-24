package internal_api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type (
	InternalApiUsecase interface {
		GetAll(c *gin.Context)
		// SIPDPS
		GetSIPDPSTanamFetch(c *gin.Context)
		GetSIPDPSProduktivitasFetch(c *gin.Context)
		GetSIPDPSPusoFetch(c *gin.Context)
		GetSIPDPSPanenFetch(c *gin.Context)
		GetSIPDPSTanam(c *gin.Context)
		GetSIPDPSProduktivitas(c *gin.Context)
		GetSIPDPSPuso(c *gin.Context)
		GetSIPDPSPanen(c *gin.Context)

		// Perbenihan
		GetPerbenihanProdusenFetch(c *gin.Context)
		GetPerbenihanRekNasFetch(c *gin.Context)
		GetPerbenihanRekBpsbFetch(c *gin.Context)
		GetPerbenihanRekLssmFetch(c *gin.Context)
		GetPerbenihanRekPenyaluranFetch(c *gin.Context)
		GetPerbenihanRekPenyebaranFetch(c *gin.Context)
		GetPerbenihanRekProdusenFetch(c *gin.Context)
		GetPerbenihanProdusen(c *gin.Context)
		GetPerbenihanRekNas(c *gin.Context)
		GetPerbenihanRekBpsb(c *gin.Context)
		GetPerbenihanRekLssm(c *gin.Context)
		GetPerbenihanRekPenyaluran(c *gin.Context)
		GetPerbenihanRekPenyebaran(c *gin.Context)
		GetPerbenihanRekProdusen(c *gin.Context)

		// SIMLUH
		GetSimluhSertifikat(c *gin.Context)
		GetSimluhRiwayatPelatihan(c *gin.Context)
	}

	usecase struct {
		repo InternalApiRepository
	}
)

func NewUsecase(repo InternalApiRepository) InternalApiUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetAll(c *gin.Context) {
	data, err := uc.repo.GetAll(c)
	if err != nil {
		// Handle the error accordingly.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetSIPDPSTanamFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_sipdps_jawa_barat")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"provinsi": "12",
			"page":     "1",
		}).
		Get("https://api-splp.layanan.go.id/t/pertanian.go.id/TP-SIPDPS/1.0/v2/data-laporan-tanam")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res SIPDPS1
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StoreSIPDPSTanamFetch(c, res.Data, 1)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetSIPDPSProduktivitasFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_sipdps_jawa_barat")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"provinsi": "12",
			"page":     "1",
		}).
		Get("https://api-splp.layanan.go.id/t/pertanian.go.id/TP-SIPDPS/1.0/v2/data-laporan-produktivitas")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res SIPDPS2
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StoreSIPDPSProduktivitasFetch(c, res.Data, 2)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetSIPDPSPusoFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_sipdps_jawa_barat")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"provinsi": "12",
			"page":     "1",
		}).
		Get("https://api-splp.layanan.go.id/t/pertanian.go.id/TP-SIPDPS/1.0/v2/data-laporan-puso")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res SIPDPS3
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StoreSIPDPSPusoFetch(c, res.Data, 3)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetSIPDPSPanenFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_sipdps_jawa_barat")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"provinsi": "12",
			"page":     "1",
		}).
		Get("https://api-splp.layanan.go.id/t/pertanian.go.id/TP-SIPDPS/1.0/v2/data-laporan-panen")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res SIPDPS4
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StoreSIPDPSPanenFetch(c, res.Data, 4)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetSIPDPSTanam(c *gin.Context) {
	data, err := uc.repo.SIPDPSTanamRead(c, 1)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetSIPDPSProduktivitas(c *gin.Context) {
	data, err := uc.repo.SIPDPSProduktivitasRead(c, 2)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetSIPDPSPuso(c *gin.Context) {
	data, err := uc.repo.SIPDPSPusoRead(c, 3)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetSIPDPSPanen(c *gin.Context) {
	data, err := uc.repo.SIPDPSPanenRead(c, 4)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanProdusenFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "seluruh",
			"jenis": "produsen",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan4
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanProdusenFetch(c, res.Data, 5)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekNasFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "nasional",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan1
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekNasFetch(c, res.Data, 6)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekBpsbFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "bpsb",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan1
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekBpsbFetch(c, res.Data, 7)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekLssmFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "lssm",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan1
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekLssmFetch(c, res.Data, 8)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekPenyaluranFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "penyaluran",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan2
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekPenyaluranFetch(c, res.Data, 9)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekPenyebaranFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "penyebaran",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan3
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekPenyebaranFetch(c, res.Data, 10)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanRekProdusenFetch(c *gin.Context) {
	token, err := uc.repo.GetToken(c, "api_token_perbenihan")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"token": token,
			"rekap": "rekapitulasi",
			"jenis": "produsen",
		}).
		Get("https://apps.tanamanpangan.pertanian.go.id/api/perbenihan")

	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Printf("error %v", response.StatusCode())
		util.JERR(c, http.StatusInternalServerError, errors.New("error"))
		return
	}

	responseBytes := response.Body()

	var res Perbenihan4
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.repo.StorePerbenihanRekProdusenFetch(c, res.Data, 11)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "success")
}

func (uc *usecase) GetPerbenihanProdusen(c *gin.Context) {
	data, err := uc.repo.PerbenihanProdusenRead(c, 5)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekNas(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekNasRead(c, 6)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekBpsb(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekBpsbRead(c, 7)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekLssm(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekLssmRead(c, 8)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekPenyaluran(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekPenyaluranRead(c, 9)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekPenyebaran(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekPenyebaranRead(c, 10)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) GetPerbenihanRekProdusen(c *gin.Context) {
	data, err := uc.repo.PerbenihanRekProdusenRead(c, 11)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

// WOKRS, DONT TOUCH V

func (uc *usecase) GetSimluhSertifikat(c *gin.Context) {
	idPel := c.DefaultQuery("id_pel", "153")
	tipe := c.DefaultQuery("tipe", "pns")
	nik := c.DefaultQuery("nik", "")

	url := "http://latihanonline.pertanian.go.id/print/print_cert_penyuluh1jt.php" +
		"?id_pel=" + idPel +
		"&nik=" + nik +
		"&tipe=" + tipe

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/pdf", resp.Body, nil)
}

func (uc *usecase) GetSimluhRiwayatPelatihan(c *gin.Context) {
	nik := c.DefaultQuery("nik", "")

	url := "https://laporanutama.pertanian.go.id/biodata/pelatihan_status_penyuluh_json.php?nik=" + nik

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var data []Pelatihan
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
