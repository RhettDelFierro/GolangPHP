package main

import (
	"net/http"
	"html/template"
	"log"
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


//Inject is pretty much where all the route handling takes place.
//Don't know if we need the http.ListenAndServe function in main.
func main() {
	//http.HandleFunc("/", inject)

	controllers.Inject(tmpl)
	http.ListenAndServe(":8080", nil)
}