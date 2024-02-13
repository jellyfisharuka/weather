package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiLey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err!=nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	err=json.Unmarshal(bytes, &c)
	if err!=nil {
		return apiConfigData{}, err
	}
	return c,nil
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", 
	func(w http.ResponseWriter, r *http.Request){
		city:= strings.SplitN(r.URL.Path, "/", 3)[2]
		data,err:=query(city)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		w.Header().Set("content-type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(data)

	})
	http.ListenAndServe(":8080", nil)
}