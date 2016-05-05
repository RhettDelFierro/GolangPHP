package main

import (
	"net/http"
	"text/template"
	"log"
	//"fmt"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/controllers"
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
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl, err = tmpl.ParseFiles("templates/_tablerows.html")
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(*tmpl)

}

//the HandlerFunc inject() was defined here.


func main() {
	//http.HandleFunc("/", inject)
	controllers.Inject(tmpl)
	http.ListenAndServe(":8080", nil)
}