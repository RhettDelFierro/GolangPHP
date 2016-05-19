package controllers

import ()
import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	//"github.com/gorilla/sessions"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"fmt"
)

type User struct {
	Data	models.UserInfo	`json:"data"`
}


//should come up with a secret key and read from a file.
//var store = sessions.NewCookieStore([]byte(""))


//there is no jwt for registerUser required.
func registerUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	user := User{}

	if req.Method == "POST" {
		//after this, you have a User{} that has a field Data
		//which is a struct of models.UserInfo{}
		//that UserInfo{} struct's fields get filled here:
		//fmt.Println("Email:", req.FormValue("email"))
		err := json.NewDecoder(req.Body).Decode(&user) //keep in mind POST from front end has a data{} object. That will go right into the field of the struct.
		//take care of writing the rest of the error later:
		if err != nil {
			//422 for json error?
			fmt.Println(err)
			fmt.Println("1st error in registerUser: json.Decode")
			//w.WriteHeader(422)
			return
		}
		//fmt.Println("Here's the data!: ", user.Data)
		userRegister := &user.Data
		fmt.Println(userRegister)
		//send to DB:
		models.RegisterUser(userRegister)
		//make sure not to send the hashed pw:
		userRegister.HashPassword = nil
		if j, err := json.Marshal(User{Data: *userRegister}); err != nil {
			fmt.Println("2nd error in registerUser. json.Marshal")
			//w.WriteHeader(422)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(j)
		}
	}
}