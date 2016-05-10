package controllers

import (
	"html/template"
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/helper"
	"strconv"
	"fmt"
)

//don't think you need a template here, you're not going to be serving the template, the javascript will manipulate the dom.
type addedController struct {
	template *template.Template
}

//errors should be custom notes from helper.
type JSON struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   map[string]map[string]string     `json:"error"`
}


//Post adds the student.
//************you still have to serve the template. But I think this is more for Populate than Add. See the getgrades.go controller.
func (this *addedController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	//initialize the struct going to json.Marshal()
	studentData := JSON{Success: false, Error: map[string]map[string]string{}}
	sd := &studentData

	//*******************maybe you can start using Go routines for all here.***************

	postMap := map[string]string{
		"name": req.FormValue("name"),
		"course": req.FormValue("course"),
		"grade": req.FormValue("grade"),
	}
	fmt.Println("here's postmap: ", postMap)
	newStudent := &helper.NewStudent{}
	//newStudent.ErrorMaker(req.FormValue("name"), req.FormValue("course"), req.FormValue("grade"), "auth_token")
	newStudent.ErrorMaker(postMap)
	fmt.Println("after ErrorMaker ", newStudent)

	regex_errors := helper.TestValidEntry(*newStudent)
	fmt.Println("here's regex_errors: ", regex_errors)
	if regex_errors != nil {
		sd.Error["regex_errors"] = regex_errors
	}

	//the regex tests will determine whether the addedstudent's info is an acceptable pattern.
	if regex_errors == nil { //include a nested if statement to check for session.
		//do the code at the bottom:
		fmt.Println("did we get to here")
		data := new(models.Student)

		grade, _ := strconv.Atoi(newStudent.Grade["value"])

		data.SetName(newStudent.Name["value"])
		data.SetCourse(newStudent.Course["value"])
		data.SetGrade(grade)
		data.SetId(bson.NewObjectId())
		convertedData := helper.StudentsToViewModel(*data)

		//don't forget to check for duplicates.
		err := models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.
		if err != nil {
			dbError := map[string]string{
				"db_error": err.Error(),
				"plain": "there was a problem inserting record into database",
			}
			sd.Error["database_error"] = dbError //helper variable for error message
		} else {
			sd.Success = true
			sd.Data = convertedData
		}

		responseWriter.Header().Add("Content-Type", "application/json")
		responseData, err := json.Marshal(sd)

		if err != nil {
			responseWriter.WriteHeader(404)
			responseWriter.Write(responseData)
		}

		responseWriter.Write(responseData)
	} else {
		//match the name of the element of the current array and return it with the "invalid" key/value.
	}

}