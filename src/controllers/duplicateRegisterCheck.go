package controllers

import (
	"net/http"
	"github.com/RhettDelFierro/GolangPHP/src/controllers/util"
	//"github.com/gorilla/sessions"
	"encoding/json"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"fmt"
	//"go/token"
	"os/user"
)

type User struct {
	Data models.UserInfo        `json:"data"`
}

type DuplicateResource struct {
	Data DuplicateModel     `json:"data"`
}

type DuplicateModel struct {
	Username    string `json:"username"`
}

type DuplicateUserInfo struct {
	Data	string	`json:"data"`
}

func DuplicateUserCheck(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	var duplicate DuplicateResource

	if req.Method == "POST"{
		err := json.NewDecoder(req.Body).Decode(&duplicate)

		if err != nil {
			//422 for json error?
			fmt.Println(err)
			fmt.Println("1st error in DuplicateUserCheck: json.Decode")
			w.WriteHeader(422)
			return
		}
		duplicateCheck := duplicate.Data
		duplicateUser := models.UserInfo{
			UserName: duplicateCheck.Username,
		}
		if userDuplicateTrue, err := models.DuplicateUser(duplicateUser); err != nil {
			//unauthorized error message
			fmt.Println("Error after DB check")
			w.WriteHeader(401)
			return
		} else {
			//user is a duplicate, generate write to response:
			fmt.Println("user.Email:", userDuplicateTrue.UserName)

			//render duplicate message on ReactJS RegisterFormContainer.
			j, err := json.Marshal(DuplicateUserInfo{Data: userDuplicateTrue.UserName})
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
}