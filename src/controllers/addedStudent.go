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
	"fmt"
	//"os"
)

type addedController struct {
	template *template.Template
}

//************you still have to serve the template. But I think this is more for Populate than Add. See the table.go controller.
func (this *addedController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	fmt.Println("logged on")
	fmt.Println("HERE IS THE NAME:", req.FormValue("name"))
	fmt.Println("HERE IS THE course:", req.FormValue("course"))
	fmt.Println("HERE IS THE grade:", req.FormValue("grade"))
	//go to the students model.
	data := new(models.Student)
	//POST
	if req.Method == "POST" { //which it will.
		data.SetName(req.FormValue("name"))
		data.SetCourse(req.FormValue("course"))
		gradeRaw := req.FormValue("grade")
		grade, err := strconv.Atoi(gradeRaw)
		if err != nil{
			panic(err)
		}
		data.SetGrade(grade)
		idRaw := req.FormValue("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil{
			panic(err)
		}
		data.SetId(id) //not going to be passed, must have the model generate it and set it. In PHP, we had mySQL increment the ID # and send it back. Mongo _id field will send it back.
	}

	convertedData := converters.StudentsToViewModel(*data)

	err := json.NewEncoder(w).Encode(convertedData)

	models.School1.AddStudents(*data) //we don't have to convert anything, just have to store it. Future videos.

	responseWriter.Header().Add("Content-Type", "application/json")
	//responseData, err := json.Marshal(convertedData)
	//fmt.Println(responseData)
	//this.template.Execute(responseWriter, as)
	if err != nil {
		responseWriter.WriteHeader(404)
	}

	//we add the students to our database above and also send it back so the front end/javascript knows we got he request.
	//responseWriter.Write(responseData)

}

//controller method to response ot the route. This is serving the data. Serve the template above, insert the data here.
//func (this addedController)