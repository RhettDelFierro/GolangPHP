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

	//the regular home page, should not load data.
	hc := new(homeController)
	hc.template = tmpl.Lookup("index.html") //may need to use the full path
	router := mux.NewRouter()
	router.HandleFunc("/index", hc.get)

	//public.
	router.HandleFunc("/api/grades", getGrades) //going to populate full student list.

	//user
	router.HandleFunc("/users/register", RegisterUser)
	router.HandleFunc("/users/login", LoginUser)
	//private
	//wrapping middleware to provide uthentication for create and delete operations.
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
	//the controllers we have for home.go and getgrades.go have no idea
	// we've used gorilla mux instead of the DefaultServerMux. The home controller doesn't need to take advantage of parameterized 	//routes, we don't have to modify them. But the CRUD routes do.



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