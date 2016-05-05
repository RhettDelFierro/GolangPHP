package controllers

import (
	"net/http"
	//"os"
	"html/template"
	//"bufio"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"os"
	"bufio"
	"fmt"
)

func Inject(tmpl *template.Template) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
		var data interface{}
		data = viewmodels.GetGrades()
		//Execute goes here.
		tmpl.ExecuteTemplate(w, "index.html", data)
	})
	http.HandleFunc("/public/", javascript)

	http.ListenAndServe(":8080", nil)
}

func javascript(w http.ResponseWriter, req *http.Request){
	path := req.URL.Path
	fmt.Println(path)

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
