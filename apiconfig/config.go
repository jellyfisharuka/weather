package apiconfig

import (
	"encoding/json"
	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

func LoadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}
