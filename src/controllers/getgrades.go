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
	//"github.com/RhettDelFierro/GolangPHP/src/controllers/converters"
	"encoding/json"
	"fmt"
	//"strconv"
	//"gopkg.in/mgo.v2/bson"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/helper"
)

type gradesController struct {
	template *template.Template
}

func (this *gradesController) getGrades(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	studentData := JSON{
		Success: false,
	}

	sd := &studentData

	//getting students from database.
	//slice of Student (not empty)
	students, err := models.GetStudents()
	if err != nil {
		//panic(err)
		sd.Error = append(studentData.Error, err.Error())
	}
	studentsVM := []viewmodels.Student{} //slice

	for _, student := range students {
		studentsVM = append(studentsVM, helper.StudentsToViewModel(student)) //have an array of hard coded Students
	}

	sd.Success = true
	sd.Data = studentsVM
	responseWriter.Header().Add("Content-Type", "application/json")
	responseData, err := json.Marshal(sd)

	//not executing a template.
	//this.template.Execute(responseWriter, responseData)

	if err != nil {
		responseWriter.WriteHeader(404)
	}

	//we add the students to our database above and
	//also send it back so the front end/javascript knows we got he request.
	responseWriter.Write(responseData)
}

//just have to look for the id.
func (this *gradesController) deleteGrade(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	fmt.Println("we got to delete")

	studentData := JSON{Success: false}
	sd := &studentData

	vars := mux.Vars(req)
	idRaw := vars["id"]

	deleted := models.DeleteStudents(idRaw)
	if (deleted) {
		sd.Success = true
		sd.Data = deleted
	} else {
		sd.Error = append(sd.Error, "entry not deleted")
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseData, err := json.Marshal(sd)
	if err != nil {
		fmt.Println("404 error", err)
		responseWriter.WriteHeader(404)
		responseWriter.Write(responseData) //will show result.error
	} else {
		responseWriter.Write(responseData)
	}

}