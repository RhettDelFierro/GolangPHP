package controllers

import(
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"text/template"
	"strconv"
	"log"
	"fmt"
)

type gradesController struct {
	template *template.Template
}

func (this *gradesController) get(w http.ResponseWriter, req *http.Request){
	id, err := strconv.Atoi(req.URL.Path[1:])
	fmt.Println(id)
	if err != nil{
		log.Fatalln(err)
	}
	vm := viewmodels.GetGrades(id)

	w.Header().Add("Content Type", "text/html")

	this.template.Execute(w, vm)
}