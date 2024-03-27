package routes

import (
	"html/template"
	"net/http"
)

var (
	tmpl1, _ = template.ParseFiles("template/template.html")
	tmpl2, _ = template.ParseFiles("template/test.html")
)

type ContactDetail struct {
	Password      string
	City          string
	Success       bool
	StorageAccess string
}
type ViewData struct {
	Title   string
	Message string
	T string
	Mes string
	Success bool
}

func Hello(w http.ResponseWriter, r *http.Request) {
	data := ContactDetail{
		Password: r.FormValue("password"),
		City:     r.FormValue("city"),
	}
	//data.Success = true
	tmpl1.Execute(w, data)
}
func Test(w http.ResponseWriter, r *http.Request) {
	testData := ViewData{
		Title:   "word",
		Message: "message",
		T: "ttt",
		Mes: "mess",
	}
   testData.Success=false
	tmpl2.Execute(w, testData)
}
