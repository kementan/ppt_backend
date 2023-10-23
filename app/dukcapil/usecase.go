package dukcapil

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IdValidation(c *gin.Context) {
	payload := []byte(`{
		"NIK": "1902011703970001",
		"NAMA_LGKP": "Ghaly Fadhillah",
		"JENIS_KLMIN": "Laki-laki",
		"TMPT_LHR": "Bandar Lampung",
		"TGL_LHR": "17-03-1997",
		"TRESHOLD": 100,
		"user_id": "26082022160454BADAN_PENYULUHAN_SDM8370",
		"password": "TH854Y",
		"ip_user": "10.160.84.10"
	}`)
	req, err := http.NewRequest("POST", "http://10.1.241.250/api/nik_verifby_elemen.php", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
		return
	}
	defer resp.Body.Close()

	var response Response

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	content := response.Content
	c.JSON(http.StatusOK, content)
}
