package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

func VerifyReCaptcha(c *gin.Context, input string) error {
	secretKey := appConfig.GSecretKey

	resp, err := http.PostForm(
		recaptchaServerName,
		url.Values{"secret": {secretKey}, "response": {input}},
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var responseData RecaptchaResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		return err
	}

	if !responseData.Success {
		return errors.New("reCAPTCHA verification failed")
	}

	return nil
}
