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
	tmpl, err = template.ParseFiles("public/dist/index.html")
	if err != nil {
		log.Fatalln(err)
	}
}

//func serveSingle(pattern string, filename string) {
//	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, filename)
//	})
//}


//the HandlerFunc inject() was defined here.


//Inject is pretty much where all the route handling takes place.
//Don't know if we need the http.ListenAndServe function in main.
func main() {
	//http.HandleFunc("/", inject)
	controllers.Inject(tmpl)
	//serveSingle("/public/dist/index_bundle.js", "./public/build/index.js")
	http.ListenAndServe(":8080", nil)
}