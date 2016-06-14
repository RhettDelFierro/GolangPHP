package controllers

import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"encoding/json"
	"fmt"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/helper"
)

type gradesController struct {
	template *template.Template
}

func getGrades(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	studentData := JSON{
		Success: false,
	}

	sd := &studentData

	//getting students from database.
	//slice of Student (not empty)
	students, err := models.GetStudents()

	//database error:
	if err != nil {
		sd.Error = append(sd.Error, err.Error()) //helper variable for error message
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(500)
		responseData, err := json.Marshal(sd)

		if err != nil {
			responseWriter.WriteHeader(404)
		}

		responseWriter.Write(responseData)

		//everything is fine:
	} else {
		studentsVM := []viewmodels.Student{} //slice

		for _, student := range students {
			studentsVM = append(studentsVM, helper.StudentsToViewModel(student))
		}

		sd.Success = true
		sd.Data = studentsVM
		responseWriter.Header().Add("Content-Type", "application/json")
		responseData, err := json.Marshal(sd)

		if err != nil {
			responseWriter.WriteHeader(404)
		}

		responseWriter.Write(responseData)

	}
}

//just have to look for the id.
func deleteGrade(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	studentData := JSON{Success: false}
	sd := &studentData

	vars := mux.Vars(req)
	idRaw := vars["id"]


	//either getting back nil or error
	deleted := models.DeleteStudents(idRaw)
	if deleted == nil {
		sd.Success = true
		sd.Data = struct{ID string}{idRaw}
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(200)
		responseData, err := json.Marshal(sd)
		if err != nil {
			fmt.Println("404 error", err)
			responseWriter.WriteHeader(404)
			responseWriter.Write(responseData)
		} else {
			responseWriter.Write(responseData)
		}
	} else { //can throw an error if the id doesn't match anything.
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(500)
		//sd.Data = deleted.ErrorID()
		sd.Error = append(sd.Error, vars["id"], deleted.Error()) //helper variable for error message
		responseData,_ := json.Marshal(sd)
		responseWriter.Write(responseData)
	}

}