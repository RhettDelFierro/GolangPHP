package controllers

import ()
import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/gorilla/sessions"
	//"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"fmt"
)

type User struct {
	Data	models.UserInfo	`json:"data"`
}


//should come up with a secret key and read from a file.
var store = sessions.NewCookieStore([]byte(""))


//there is no jwt for registerUser required.
func registerUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	user := User{}

	if req.Method == "POST" {
		//after this, you have a User{} that has a field Data
		//which is a struct of models.UserInfo{}
		//that UserInfo{} struct's fields get filled here:
		err := json.NewDecoder(req.Body).Decode(&user)
		//take care of writing the rest of the error later
		if err != nil {
			//422 for json error?
			fmt.Println("1st error in registerUser: json.Decode")
			w.WriteHeader(422)
		}
		userRegister := &user.Data

		//send to DB:
		models.RegisterUser(userRegister)
		//make sure not to send the hashed pw:
		userRegister.HashPassword = nil
		if j, err := json.Marshal(User{Data: *userRegister}); err != nil {
			fmt.Println("2nd error in registerUser. json.Marshal")
			w.WriteHeader(422)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(j)
		}
	}
}

func loginUser(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		//422 for json error?
		w.WriteHeader(422)
	}

	//make call to DB to get user info.

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