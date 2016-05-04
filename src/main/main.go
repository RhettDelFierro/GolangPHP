package main

import (
	"net/http"
	"html/template"
	"log"
	//"fmt"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
)

type Students struct{
	Name 	string
	Course 	string
	Grade 	int
}

type School struct{
	All []Students
}

var tmpl *template.Template

//parse the html template files
func init() {
	var err error
	tmpl, err = template.ParseFiles("../../templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl, err = tmpl.ParseFiles("../../templates/_tablerows.html")
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(*tmpl)

}

//HandlerFunc
func inject(w http.ResponseWriter, req *http.Request){
	var data interface{}
	data = viewmodels.GetGrades()
	//Execute goes here.
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	http.HandleFunc("/", inject)
	http.ListenAndServe(":8080", nil)
}