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
	"github.com/gorilla/mux"
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

	//creating a router to takes the responsibilities for delivering our pages away from us (us = Golang's net/http std library)
	router := mux.NewRouter() //assigning the variable "router" to the object returned by a call to gorilla mux's NewRouter() function. This is just like the Handler interface now. Just need it for our two pages, because the HandleFunc for scripts doesn't need any parameterized routing.

	//the regular home page, should not load data.
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path
	router.HandleFunc("/index", hc.get)

	//on click go to this page (but right now don't have a link to it in html files)
	//gc := new(gradesController)
	//gc.template = tmpl.Lookup("index.html") //index.html here because it is set to include the data injected.
	//router.HandleFunc("/grades", gc.get)

	//now we have to set the net/http package to set the gorilla mux router (variable "router") to listen for requests.
	http.Handle("/", router) //the controllers we have for home.go and table.go have no idea we've used gorilla mux instead of the DefaultServerMux. The home controller doesn't need to take advantage of parameterized routes, we don't have to modify them. But the table controller does.
	//creating a category controller and register it with the router
	gradesController := new(gradesController)
	gradesController.template = tmpl.Lookup("index.html")
	//fmt.Println(gradesController.template)
	router.HandleFunc("/grades/{id}", gradesController.get) //the {id} curly braces is how we indicate to gorilla mux that we want to grab this part of the route path and map it to the "id" key in the route map.


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
