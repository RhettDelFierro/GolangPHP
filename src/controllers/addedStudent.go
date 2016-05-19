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

//errors should be custom notes from helper.
type JSON struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
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

	//*******************maybe you can start using Go routines for all here.***************

	postMap := map[string]string{
		"name": req.FormValue("name"),
		"course": req.FormValue("course"),
		"grade": req.FormValue("grade"),
	}

	fmt.Println(postMap)
	//fmt.Println("here's postmap: ", postMap)
	//newStudent := &helper.NewStudent{}
	//newStudent.ErrorMaker(req.FormValue("name"), req.FormValue("course"), req.FormValue("grade"), "auth_token")
	regexCheckingMap := helper.ErrorMaker(postMap)
	//fmt.Println("after ErrorMaker ", regexCheckingMap)


	regex_errors := helper.TestValidEntry(regexCheckingMap)
	fmt.Println("here's regex_errors: ", regex_errors)
	if regex_errors != nil {

		for i, _ := range regex_errors {
			sd.Error = append(sd.Error, regex_errors[i])
		}
	}

	//the regex tests will determine whether the addedstudent's info is an acceptable pattern.
	if len(regex_errors) == 0 {
		//include a nested if statement to check for session.
		//do the code at the bottom:
		fmt.Println("did we get to here")
		data := new(models.Student)

		grade, _ := strconv.Atoi(regexCheckingMap["grade"]["value"])

		data.SetName(regexCheckingMap["name"]["value"])
		data.SetCourse(regexCheckingMap["course"]["value"])
		data.SetGrade(grade)
		data.SetId(bson.NewObjectId())
		convertedData := helper.StudentsToViewModel(*data)

		//don't forget to check for duplicates.
		duplicate, err := models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.
		//checking to see for either an error in the DB or duplicate.
		if err != nil || len(duplicate) != 0 {
			fmt.Println("here is an error for err: ", err)
			fmt.Println("here is what duplicate is printing", duplicate)
			//if error in DB connection
			if err != nil {
				fmt.Println("we should be in here now")
				sd.Error = append(sd.Error, err.Error())
				responseWriter.Header().Add("Content-Type", "application/json")
				responseWriter.WriteHeader(500)
				responseData, err := json.Marshal(sd)
				if err != nil {
					fmt.Println("error in json writing")
					responseWriter.Write(responseData)
				} else {
					//sd.Error = append(sd.Error, err.Error())
					fmt.Println("no error in json writing")
					responseWriter.Write(responseData)
				}
			}
			//if no error in DB connection, but duplicate student:
			if len(duplicate) != 0 {
				//should make this a function because getgrades.go uses the same thing. DRY
				fmt.Println(duplicate)
				sd.Error = append(sd.Error, regexCheckingMap["duplicate"]["error"])
				students := duplicate
				studentsVM := []viewmodels.Student{} //slice

				for _, student := range students {
					studentsVM = append(studentsVM, helper.StudentsToViewModel(student))
				}
				sd.Data = studentsVM
				fmt.Printf(" type of sd.Data %T \n", sd.Data)
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