package controllers

import (
	"net/http"
	"html/template"
	"os"
	"bufio"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

//this is pretty much where all the route handlers are.
func Inject(tmpl *template.Template) {

	//I think we can kill the templating?

	//dispensing the ReactJS
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path

	router := mux.NewRouter()
	router.HandleFunc("/", hc.get)

	//public.
	router.HandleFunc("/api/grades", getGrades) //going to populate full student list.

	//check for duplicate usernames:
	router.HandleFunc("/username", DuplicateNewUserCheck)

	//user
	router.HandleFunc("/users/register", RegisterUser)
	router.HandleFunc("/users/login", LoginUser)

	//private
	//wrapping middleware to provide authentication for create and delete operations.
	router.PathPrefix("/api/delete/{id}").Handler(
		negroni.New(
			negroni.HandlerFunc(AuthorizeToken),
			negroni.Wrap(http.HandlerFunc(deleteGrade))))
	router.PathPrefix("/api/add").Handler(
		negroni.New(
		negroni.HandlerFunc(AuthorizeToken),
		negroni.Wrap(http.HandlerFunc(postStudent))))


	//necessary stuff. To set up the above.
	//now we have to set the net/http package to set the gorilla mux router
	//(variable "router") to listen for requests.
	http.Handle("/", router)


	//do we need this? The first line, yes.
	http.HandleFunc("/index_bundle.js", javascript)
	//http.ListenAndServe(":8080", nil)
}

func javascript(w http.ResponseWriter, req *http.Request){
	path := "public/dist/" + req.URL.Path

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