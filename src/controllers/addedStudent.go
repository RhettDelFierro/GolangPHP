package controllers

import (
	"html/template"
	"net/http"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"strconv"
	//"fmt"
	//"os"
	"gopkg.in/mgo.v2/bson"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"fmt"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/helper"
)

//don't think you need a template here, you're not going to be serving the template, the javascript will manipulate the dom.
type addedController struct {
	template *template.Template
}

type JSON struct {
	Success	bool	`json:"success"`
	Data interface{} `json:"data"`
	Error []string	`json:"error"`
}


//************you still have to serve the template. But I think this is more for Populate than Add. See the getgrades.go controller.
func (this *addedController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	studentData := JSON{Success: false}

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
	//if (convertedData) {
	//	studentData.Success = true
	//	studentData.Data = convertedData
	//} else {
	//	studentData.Error = append(studentData.Error, "database error")
	//}
	//err := json.NewEncoder(w).Encode(convertedData)
	fmt.Println("convertedData: ", convertedData)
	err := models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.
	if err != nil {
		studentData.Success = false
		studentData.Error = append(studentData.Error, err.Error()) //maybe just append to error array "student not added"
		//responseWriter.Write
	} else {
		studentData.Success = true
		studentData.Data = convertedData
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseData, err := json.Marshal(studentData)
	fmt.Println("here is the converted JSON data:", studentData)
	//not executing a template.
	//this.template.Execute(responseWriter, responseData)
	if err != nil {
		responseWriter.WriteHeader(404) //result.error on the front end.
		//responseWriter.Write(err.Error()) //write the result.error on front end.
	}

	//we add the students to our database above and also send it back so the front end/javascript knows we got he request.
	responseWriter.Write(responseData) //"result" on the front end. Write the errors also, do something with them on the front end. The
	fmt.Println(helper.AddingStudent)
}

//controller method to response ot the route. This is serving the data. Serve the template above, insert the data here.
//func (this addedController)