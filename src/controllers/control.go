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
	router := mux.NewRouter()

	//I think we can kill the templating?

	//the regular home page, should not load data.
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path
	router.HandleFunc("/index", hc.get)

	//on click go to this page (but right now don't have a link to it in html files)
	//gc := new(addedController) //gradesController is package level so you're good.
	//gc.template = tmpl.Lookup("index.html") //index.html here because it is set to include the data injected.
	router.HandleFunc("/api/add", addedController.post) //anything that goes to /grades will be handled by ajaxMethods.

	//now we have to set the net/http package to set the gorilla mux router (variable "router") to listen for requests.
	http.Handle("/", router) //the controllers we have for home.go and table.go have no idea we've used gorilla mux instead of the DefaultServerMux. The home controller doesn't need to take advantage of parameterized routes, we don't have to modify them. But the table controller does.

	//creating a grades controller and register it with the router
	gradesController := new(gradesController)
	gradesController.template = tmpl.Lookup("index.html")
	//fmt.Println(gradesController.template)
	router.HandleFunc("/grades/{id}", gradesController.ajaxMethods) //the {id} curly braces is how we indicate to gorilla mux that we want to grab this part of the route path and map it to the "id" key in the route map.


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