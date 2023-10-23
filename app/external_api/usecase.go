package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/config"
	"github.com/gin-gonic/gin"
)

var (
	appConfig, _ = config.LoadConfig("./.")
)

type (
	ExternalApiUsecase interface {
		GetWeather(c *gin.Context)
	}

	usecase struct{}
)

func NewUsecase() ExternalApiUsecase {
	return &usecase{}
}

func (uc *usecase) GetWeather(c *gin.Context) {
	lat := c.DefaultQuery("lat", "-6.305964")
	lng := c.DefaultQuery("lng", "106.819983")
	if lat == "" || lng == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and Longitude are required"})
		return
	}

	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lng + "&appid=" + appConfig.OWM_API

	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	main := data["main"].(map[string]interface{})
	tempKelvin := main["temp"].(float64)
	tempCelsius := tempKelvin - 273.15
	formattedTemp := fmt.Sprintf("%.1f", tempCelsius)

	weather := data["weather"].([]interface{})
	weatherData := weather[0].(map[string]interface{})
	weatherDescription := weatherData["description"].(string)
	weatherMain := weatherData["main"].(string)
	weatherIcon := weatherData["icon"].(string)

	humidity := main["humidity"].(float64)
	pressure := main["pressure"].(float64)

	wind := data["wind"].(map[string]interface{})
	windDirection := wind["deg"].(float64)
	windSpeed := wind["speed"].(float64)

	result := gin.H{
		"temp_celcius":   formattedTemp,
		"weather":        weatherMain,
		"weather_desc":   weatherDescription,
		"weather_icon":   weatherIcon,
		"humidity":       humidity,
		"pressure":       pressure,
		"wind_direction": windDirection,
		"wind_speed":     windSpeed,
	}

	c.JSON(http.StatusOK, result)
}
