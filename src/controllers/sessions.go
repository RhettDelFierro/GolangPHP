package controllers

import ()
import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/gorilla/sessions"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("flabblegabble"))

func loginUser(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	//if they submitted a form
	//if they don't it'll be a GET request
	//and the code in the if block won't apply:
	if req.Method == "POST" {
		//we're getting the session out of the user's cookies (because it's stored there):
		//we're decrypting the cookie.
		session, _ := store.Get(req, "session-name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		//check to make sure they're a valid user:
		if email == "password from DB" && password == "some password from DB" {
			//if they are,
			//set the session value.
			session.Values["logged_in"] = "this user is logged in"
		} else {
			//just do the sd struct again.
			http.Error(w, "invalid credentials", 400)
			return
		}
		session.Save(req, w)

		//if they're all valid, redirect them to user stuff:
		//third parameter is the url you want to redirect them to:
		//we're going to send them to the main page. Which will then have a
		//user logged in bar at the top.
		//302 is a common redirection code.
		http.Redirect(w, req, "/", 302)
		return
	}

	//this is the code if we didn't do any POST stuff:
	//we're just seeing the form normally:
	//we're just going to be rendering the form templates:
}

func logoutUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	session, _ := store.Get(req, "session-name")

	delete(session.Values, "logged_in")
	session.Save(req, w)
	http.Redirect(w, req, "/", 302)
}


//this is just a general overall example and includes all situations. The real applications are above.
func sessionRoutes(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	//If no sessions exists, this also creates a brand new one.
	session, err := store.Get(req, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//once we have the session:

	//get a value:
	//if it wasn't already in there, it will be nil.
	value := session.Values["whatever_value_it_is_I_am_looking_for"]

	//to get this as a string and not an interface{}
	//string, _ := session.Values["whatever_value_it_is_I_amd_looking_for"].(string)

	//maybe even:
	str, _ := value.(string)
	//in general using strings makes it easier.
	fmt.Println(str)

	// Set some session values.
	//the session type is map[interface{}]interface{} which means
	//the map keys and values can be of any type.
	session.Values["whatever_value_it_is_I_am_looking_for"] = "98372459874328957"
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	//delete a value:
	//this will remove the session.
	delete(session.Values, "whatever_value_it_is_I_amd_looking_for")
	// Save it before we write to the response/return from the handler.
	session.Save(req, w)
}
