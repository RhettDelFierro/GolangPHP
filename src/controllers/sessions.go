package controllers

import ()
import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("flabblegabble"))

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
	value := session.Values["whatever_value_it_is_I_amd_looking_for"]

	//to get this as a string and not an interface{}
	//string, _ := session.Values["whatever_value_it_is_I_amd_looking_for"].(string)

	//maybe even:
	string, _ := value.(string)

	// Set some session values.
	//the session type is map[interface{}]interface{} which means
	//the map keys and values can be of any type.
	session.Values["whatever_value_it_is_I_amd_looking_for"] = 98372459874328957
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	//delete a value:
	//this will remove the session.
	delete(session.Values, "whatever_value_it_is_I_amd_looking_for")
	// Save it before we write to the response/return from the handler.
	session.Save(req, w)
}
