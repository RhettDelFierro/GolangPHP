package controllers

import ()
import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/gorilla/sessions"
	//"fmt"
)

//should come up with a secret key and read from a file.
var store = sessions.NewCookieStore([]byte(""))

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