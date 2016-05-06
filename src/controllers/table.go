package controllers

import(
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"text/template"
	"strconv"
	//"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
)

type gradesController struct {
	template *template.Template
}

func (this *gradesController) get(w http.ResponseWriter, req *http.Request){
	//id, err := strconv.Atoi(req.URL.Path[1:])
	//fmt.Println(id)
	//if err != nil{
	//	log.Fatalln(err)
	//}

	//gorillamux is being used here
		//we need to get the parameter container that gorilla mux stores for each request.
	vars := mux.Vars(req) //returns a map of the parameters caught by the current request
	idRaw := vars["id"] //be aware mux provides the values as a string, but for viewmodels.GetGrades() we need an integer for the argument.
	id, err := strconv.Atoi(idRaw)
	fmt.Println(id)
	if err !=nil {
		w.WriteHeader(404)
		panic(err)
	}

	vm := viewmodels.GetGrades(id)
	w.Header().Add("Content Type", "text/html")
	//fmt.Println(vm)

	//this.template.Execute(w, vm)
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}