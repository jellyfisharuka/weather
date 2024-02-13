package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"weather/routes"

)


func main() {
	http.HandleFunc("/hello", routes.Hello)
	http.HandleFunc("/weather/", 
	func(w http.ResponseWriter, r *http.Request){
		city:= strings.SplitN(r.URL.Path, "/", 3)[2]
		data,err:=routes.Query(city)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		w.Header().Set("content-type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(data)

	})
	http.ListenAndServe(":8080", nil)
}