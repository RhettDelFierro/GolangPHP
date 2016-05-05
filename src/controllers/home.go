package controllers

import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request){
	vm := viewmodels.GetGrades()

	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}