package controllers

import(
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"text/template"
	"strconv"
	//"log"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/converters"
)

type gradesController struct {
	template *template.Template
}

func (this *gradesController) ajaxMethods(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw) //id is mainly for delete.

	if err !=nil {
		w.WriteHeader(404)
		panic(err)
	}





	students := models.GetStudents() //slide of Student (not empty)
	studentsVM := []viewmodels.Students{} //slice

	for _, student := range students {
		studentsVM = append(studentsVM, converters.StudentsToViewModel(student)) //have an array of hard coded Students
	}

	vm := viewmodels.GetStudents()
	vm.Students = studentsVM //now we have a view model Struct with field Students that is an array of viewmodel.Students Schools.Students. Ready to execute/inject the view with these, as seen below. May now have to work the HTML injecting to use range.

	responseWriter.Header().Add("Content Type", "text/html")

	//Create
	this.template.Execute(responseWriter, vm)
}