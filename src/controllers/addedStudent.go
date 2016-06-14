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
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
)

//don't think you need a template here, you're not going to be serving the template, the javascript will manipulate the dom.
type addedController struct {
	template *template.Template
}

type StudentJSON struct {
	Name string `json:"name"`
	Course string `json:"course"`
	Grade string `json:"grade"`
}

//errors should be custom notes from helper.
type JSON struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"student"`
	Error   []string     `json:"error"`
}


//Post adds the student.
//************you still have to serve the template. But I think this is more for Populate than Add. See the getgrades.go controller.
func postStudent(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	//initialize the struct going to json.Marshal()
	studentData := JSON{Success: false}
	sd := &studentData

	student := StudentJSON{}
	err := json.NewDecoder(req.Body).Decode(&student)
	fmt.Println(student, err)

	//*******************maybe you can start using Go routines for all here.***************

	postMap := map[string]string{
		"name": student.Name,
		"course": student.Course,
		"grade": student.Grade,
	}

	regexCheckingMap := helper.ErrorMaker(postMap)
	regex_errors := helper.TestValidEntry(regexCheckingMap)

	if regex_errors != nil {
		for i, _ := range regex_errors {
			sd.Error = append(sd.Error, regex_errors[i])
		}
	}

	//the regex tests will determine whether the addedstudent's info is an acceptable pattern.
	if len(regex_errors) == 0 {
		data := new(models.Student)

		grade, _ := strconv.Atoi(student.Grade)

		data.SetName(student.Name)
		data.SetCourse(student.Course)
		data.SetGrade(grade)
		data.SetId(bson.NewObjectId())
		convertedData := helper.StudentsToViewModel(*data)

		duplicate, err := models.AddStudents(data)
		//checking to see for either an error in the DB or duplicate.
		if err != nil || len(duplicate) != 0 {
			//if error in DB connection
			if err != nil {
				sd.Error = append(sd.Error, err.Error())
				responseWriter.Header().Add("Content-Type", "application/json")
				responseWriter.WriteHeader(500)
				responseData, err := json.Marshal(sd)
				if err != nil {
					responseWriter.Write(responseData)
				} else {
					//sd.Error = append(sd.Error, err.Error())
					responseWriter.Write(responseData)
				}
			}
			//if no error in DB connection, but duplicate student:
			if len(duplicate) != 0 {
				//should make this a function because getgrades.go uses the same thing. DRY
				sd.Error = append(sd.Error, regexCheckingMap["duplicate"]["error"])
				students := duplicate
				studentsVM := []viewmodels.Student{} //slice

				for _, student := range students {
					studentsVM = append(studentsVM, helper.StudentsToViewModel(student))
				}
				sd.Data = studentsVM
				responseWriter.Header().Add("Content-Type", "application/json")
				responseWriter.WriteHeader(400)
				responseData, err := json.Marshal(sd)
				if err != nil {
					panic(err)
				} else {
					responseWriter.Write(responseData)
				}


			}








		//everything checks out
		} else {
			sd.Success = true
			sd.Data = convertedData
		}
		if len(sd.Error) == 0 {
			responseWriter.Header().Add("Content-Type", "application/json")
			responseWriter.WriteHeader(200)
			responseData, err := json.Marshal(sd)

			//error writing the JSON
			if err != nil {
				//responseWriter.WriteHeader(500)
				responseWriter.Write(responseData)
			}

			responseWriter.Write(responseData)
		}
		//regex doesn't pass
	} else {
		//I think need to throw in 500 errors for the json marshalling errors
		//match the name of the element of the current array and return it with the "invalid" key/value.
		responseWriter.WriteHeader(400)
		responseData, err := json.Marshal(sd)

		if err != nil {
			//responseWriter.WriteHeader(500)
			responseWriter.Write(responseData)
		}

		responseWriter.Write(responseData)
	}

}