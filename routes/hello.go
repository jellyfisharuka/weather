package routes

import (
	"net/http"
	"text/template"
)

var (
	tmpl = template.Must(template.ParseFiles("C:/Users/aruke/OneDrive/Рабочий стол/weather/routes/template/template.html"))
)
type ContactDetail struct {
	Password string
	Success bool
	City string
	StorageAccess string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go!\n"))
	w.Header().Set("Content-Type", "text/html") 
	http.ServeFile(w, r, "C:/Users/aruke/OneDrive/Рабочий стол/weather/routes/template/template.html")

	data := ContactDetail{
		City:     r.FormValue("city"),
		Password: r.FormValue("password"),
	}
	tmpl.Execute(w, data)
}
