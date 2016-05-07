package controllers

import(
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"html/template"
	//"strconv"
	//"log"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/converters"
	"encoding/json"
	"fmt"
	//"strconv"
	"gopkg.in/mgo.v2/bson"
)

type gradesController struct {
	template *template.Template
}

type ResponseJSON struct {
	body []byte
}

func (this *gradesController) getGrades(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	fmt.Println("ajaxMethods called")

	students, err := models.GetStudents() //slice of Student (not empty)
	if err != nil {
		panic(err)
	}
	studentsVM := []viewmodels.Student{} //slice

	for _, student := range students {
		studentsVM = append(studentsVM, converters.StudentsToViewModel(student)) //have an array of hard coded Students
	}

	vm := viewmodels.GetStudents()
	vm.Students = studentsVM //now we have a view model Struct with field Students that is an array of viewmodel.Students Schools.Students. Ready to execute/inject the view with these, as seen below. May now have to work the HTML injecting to use range.

	//responseWriter.Header().Add("Content Type", "text/html")

	//Create
	//this.template.Execute(responseWriter, vm)

	//Expose the fields from data *models.Student
	//convertedData := converters.StudentsToViewModel(*data)

	//err := json.NewEncoder(w).Encode(convertedData)

	//models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.
	//fmt.Println(vm.Students)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseData, err := json.Marshal(vm.Students)
	objectJSON := ResponseJSON{body: responseData}

	//not executing a template.
	//this.template.Execute(responseWriter, responseData)
	if err != nil {
		responseWriter.WriteHeader(404)
	}

	//we add the students to our database above and also send it back so the front end/javascript knows we got he request.
	responseWriter.Write(objectJSON.body)
}

//just have to look for the id.
func (this *gradesController) deleteGrade(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req)
	idRaw := vars["id"]
	id := bson.ObjectId(idRaw) //id is mainly for delete.
	success := models.DeleteStudents(id)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseData, err := json.Marshal(success)
	objectJSON := ResponseJSON{body: responseData}
	if err != nil {
		responseWriter.WriteHeader(404)
	} else {
		responseWriter.Write(objectJSON.body)
	}



}