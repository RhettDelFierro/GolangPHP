package main

import (
	"net/http"
	"html/template"
	"log"
)

type Students struct{
	Name 	string
	Course 	string
	Grade 	int
}

var tmpl *template.Template

//parse the html template file
func init() {
	var err error
	tmpl, err = template.ParseFiles("../../templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}

}

//HandlerFunc
func inject(w http.ResponseWriter, req *http.Request){
	student1 := Students{Name: "Student1", Course: "Course1", Grade: 100}
	//Execute goes here.
	tmpl.Execute(w, student1)
}

func main() {
	http.HandleFunc("/", inject)
	http.ListenAndServe(":8080", nil)
}