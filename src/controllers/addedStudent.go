package controllers

import (
	"html/template"
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/helper"
)

//don't think you need a template here, you're not going to be serving the template, the javascript will manipulate the dom.
type addedController struct {
	template *template.Template
}

//errors should be custom notes from helper.
type JSON struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   []string        `json:"error"`
}


//Post adds the student.
//************you still have to serve the template. But I think this is more for Populate than Add. See the getgrades.go controller.
func (this *addedController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	//*******************maybe you can start using Go routines for all here.***************

	postMap := map[string]string{
		"name": req.FormValue("name"),
		"course": req.FormValue("course"),
		"grade": req.FormValue("grade"),
	}

	newStudent := &helper.NewStudent{}
	//newStudent.ErrorMaker(req.FormValue("name"), req.FormValue("course"), req.FormValue("grade"), "auth_token")
	newStudent.ErrorMaker(postMap)

	catchRegexArray := helper.TestValidEntry(*newStudent)

	//the regex tests will determine whether the addedstudent's info is an acceptable pattern.
	if (len(catchRegexArray) == 0){
		//do the code at the bottom:
	} else {
		//match the name of the element of the current array and return it with the "invalid" key/value.
	}

	fmt.Println(newStudent)

	studentData := JSON{Success: false}
	sd := &studentData

	//go to the students model.
	data := new(models.Student)
	//POST
	if req.Method == "POST" {
		//which it will.
		data.SetName(req.FormValue("name"))
		data.SetCourse(req.FormValue("course"))
		gradeRaw := req.FormValue("grade")
		grade, err := strconv.Atoi(gradeRaw)
		if err != nil {
			panic(err)
		}
		data.SetGrade(grade)
		data.SetId(bson.NewObjectId())
	}

	//Expose the fields from data *models.Student otherwise it won't be seen
	convertedData := helper.StudentsToViewModel(*data)

	err := models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.
	if err != nil {
		sd.Success = false
		sd.Error = append(sd.Error, err.Error()) //helper variable for error message
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
}

//controller method to response ot the route. This is serving the data. Serve the template above, insert the data here.
//func (this addedController)