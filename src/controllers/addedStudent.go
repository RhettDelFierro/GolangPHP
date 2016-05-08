package controllers

import (
	"html/template"
	"net/http"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/converters"
	"strconv"
	//"fmt"
	//"os"
	"gopkg.in/mgo.v2/bson"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"fmt"
)

//don't think you need a template here, you're not going to be serving the template, the javascript will manipulate the dom.
type addedController struct {
	template *template.Template
}

type JSONAdd struct {
	Data viewmodels.Student `json:"data"`
}

//************you still have to serve the template. But I think this is more for Populate than Add. See the getgrades.go controller.
func (this *addedController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

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
	convertedData := converters.StudentsToViewModel(*data)
	studentData := JSONAdd{Data: convertedData}
	//err := json.NewEncoder(w).Encode(convertedData)
	fmt.Println("convertedData: ", convertedData)
	models.AddStudents(data) //we don't have to convert anything, just have to store it. Future videos.

	responseWriter.Header().Add("Content-Type", "application/json")
	//responseData, err := json.Marshal(convertedData)
	responseData, err := json.Marshal(studentData)
	fmt.Println("here is the converted JSON data:", studentData)
	//not executing a template.
	//this.template.Execute(responseWriter, responseData)
	if err != nil {
		responseWriter.WriteHeader(404)
	}

	//we add the students to our database above and also send it back so the front end/javascript knows we got he request.
	responseWriter.Write(responseData)

}

//controller method to response ot the route. This is serving the data. Serve the template above, insert the data here.
//func (this addedController)