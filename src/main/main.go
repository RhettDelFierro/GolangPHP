package main

import (
	"net/http"
	"html/template"
	"log"
	"fmt"
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
	student1 := Students{Name: "Student1", Course: "Course1", Grade: 100,}
	student2 := Students{Name: "Student2", Course: "Course2", Grade: 100,}
	student3 := Students{Name: "Student1", Course: "Course3", Grade: 100,}
	class := School{All: []Students{student1, student2, student3}}
	fmt.Println(class)
	//Execute goes here.
	tmpl.Execute(w, class)
}

func main() {
	http.HandleFunc("/", inject)
	http.ListenAndServe(":8080", nil)
}