package controllers

import (
	"net/http"
	//"os"
	"html/template"
	//"bufio"
	//"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"os"
	"bufio"
	//"fmt"
	"github.com/gorilla/mux"
	//"fmt"
)

//this is pretty much where all the route handlers are.
func Inject(tmpl *template.Template) {
	router := mux.NewRouter()

	//I think we can kill the templating?

	//the regular home page, should not load data.
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path
	router.HandleFunc("/index", hc.get)

	//addStudents.
	//ac := new(addedController) //gradesController is package level so you're good.
	//ac.template = tmpl.Lookup("index.html") //index.html here because it is set to include the data injected.
	router.HandleFunc("/api/add", postStudent) //anything that goes to /grades will be handled by ajaxMethods.


	gradesController := new(gradesController)
	router.HandleFunc("/api/grades", gradesController.getGrades) //going to populate full student list.
	router.HandleFunc("/api/delete/{id}", gradesController.deleteGrade)

	//session
	router.HandleFunc("/users/login", loginUser)
	router.HandleFunc("/users/logout", logoutUser)

	router.HandleFunc("/users/register", registerUser)




	//necessary stuff. To set up the above.
	//now we have to set the net/http package to set the gorilla mux router (variable "router") to listen for requests.
	http.Handle("/", router) //the controllers we have for home.go and getgrades.go have no idea we've used gorilla mux instead of the DefaultServerMux. The home controller doesn't need to take advantage of parameterized routes, we don't have to modify them. But the grades controller does.
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