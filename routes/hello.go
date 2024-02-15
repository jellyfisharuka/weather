package routes

import (
	"html/template"
	"net/http"
)

var (
	tmpl = template.Must(template.ParseFiles("template/template.html"))
)

type ContactDetail struct {
	Password      string
	Success       bool
	City          string
	StorageAccess string
}
type ViewData struct {
	Title   string
	Message string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go!\n"))
	w.Header().Set("Content-Type", "html/html")

	data := ContactDetail{
		City:     r.FormValue("city"),
		Password: r.FormValue("password"),
	}
	tmpl.Execute(w, data)
}
func Test(w http.ResponseWriter, r *http.Request) {
	testData := ViewData{
		Title:   "word",
		Message: "message",
	}
	tmpl, _ := template.ParseFiles("template/test.html")
	tmpl.Execute(w, testData)
}
