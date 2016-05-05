package controllers

import (
	"net/http"
	//"os"
	"text/template"
	//"bufio"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"os"
	"bufio"
	//"fmt"
)

func Inject(tmpl *template.Template) {
	//http.HandleFunc("/",
	//	func(w http.ResponseWriter, req *http.Request) {
	//	var data interface{}
	//	data = viewmodels.GetGrades()
	//	//Execute goes here.
	//	tmpl.ExecuteTemplate(w, "index.html", data)
	//})

	//te regular home page, should not load data.
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path
	http.HandleFunc("/index", hc.get)

	//on click go to this page (but right now don't have a link to it in html files)
	gc := new(gradesController)
	gc.template = tmpl.Lookup("index.html") //index.html here because it is set to include the data injected.
	http.HandleFunc("/grades", gc.get)

	http.HandleFunc("/scripts/", javascript)

	http.ListenAndServe(":8080", nil)
}

func javascript(w http.ResponseWriter, req *http.Request){
	path := "public" + req.URL.Path

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()

		w.Header().Add("Content Type", "text/javascript")
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
