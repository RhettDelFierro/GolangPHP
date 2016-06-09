package controllers

import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	//"github.com/gorilla/sessions"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"fmt"
	//"go/token"
)

type User struct {
	Data models.UserInfo        `json:"data"`
}

type LoginResource struct {
	Data LoginModel        `json:"data"`
}

type LoginModel struct {
	User	string	`json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthToken struct {
	User  models.UserInfo `json:"user"`
	Token string                `json:"token"`
}

type AuthTokenSent struct {
	Data AuthToken        `json:"data"`
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

		//send to DB:
		models.RegisterUser(userRegister)
		//make sure not to send the hashed pw:
		userRegister.HashPassword = nil
		if j, err := json.Marshal(userRegister); err != nil {
			w.WriteHeader(422)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(j)
		}
	}
}

func LoginUser(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	//login will be the info taken from login form
	var login LoginResource
	//token for the jwt token once authentication from login
	//is confirmed.
	var token string

	if req.Method == "POST" {
		err := json.NewDecoder(req.Body).Decode(&login)
		if err != nil {
			//422 for json error?
			fmt.Println(err)
			fmt.Println("1st error in LoginUser: json.Decode")
			w.WriteHeader(422)
			return
		}
	}
	//authenticate user
	loginModel := login.Data
	loginUser := models.UserInfo{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	//database check:
	if user, err := models.CheckUser(loginUser); err != nil {
		//unauthorized error message
		fmt.Println("Error after DB check")
		w.WriteHeader(401)
		return
	} else {
		//user is verified, generate jwt:
		fmt.Println("user.Email:", user.Email)
		token, err = GenerateToken(user.Email, "teacher")
		if err != nil {
			fmt.Println("Error generating token")
			w.WriteHeader(500)
			w.Write([]byte("Could not generate token"))
			return
		}

		//don't send the HashPassword.
		fmt.Println("heres the token:", token)
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		//send this token and code to the front end.
		authUser := AuthToken{
			User: user,
			Token: token,

		}

		//taken the response on the front end and throw it in the header.
		//include the header for all Update/Add/Delete CRUD methods.
		j, err := json.Marshal(AuthTokenSent{Data: authUser})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("An unexpected error has occured. Json not wrote."))
			return
		} else {
			w.WriteHeader(200)
			w.Write(j)
		}
	}
}