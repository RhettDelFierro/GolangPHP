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

type LoginResource struct {
	Data	LoginModel	`json:"data"`
}

type LoginModel struct {
	Email		string `json:"email"`
	Password	string `json:"password"`
}

//there is no jwt for registerUser required.
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	user := User{}

	if req.Method == "POST" {
		//after this, you have a User{} that has a field Data
		//which is a struct of models.UserInfo{}
		//that UserInfo{} struct's fields get filled here:
		err := json.NewDecoder(req.Body).Decode(&user) //keep in mind POST from front end has a data{} object. That will go right into the field of the struct.
		//take care of writing the rest of the error later:
		if err != nil {
			//422 for json error?
			fmt.Println(err)
			fmt.Println("1st error in registerUser: json.Decode")
			w.WriteHeader(422)
			return
		}
		userRegister := &user.Data
		fmt.Println(userRegister)
		//send to DB:
		models.RegisterUser(userRegister)
		//make sure not to send the hashed pw:
		userRegister.HashPassword = nil
		if j, err := json.Marshal(User{Data: *userRegister}); err != nil {
			w.WriteHeader(422)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(j)
		}
	}
}

func LoginUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	var login LoginResource
	var token string

	if req.Method == "POST" {
		err := json.Decoder(req.Body).Decode(&login)
		if err != nil {
			//422 for json error?
			fmt.Println(err)
			fmt.Println("1st error in registerUser: json.Decode")
			w.WriteHeader(422)
			return
		}
	}

	loginModel := login.Data
	loginUser := models.UserInfo{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	if user, err := models.CheckUser(loginUser); err != nil {
		//unauthorized error message
		w.WriteHeader(401)
		return
	} else {
		token, err = GenerateToken(user.Email, "teacher")
		if err != nil {
			w.WriteHeader(500)
			w.Write("Could not generate token")
			return
		}
	}
}