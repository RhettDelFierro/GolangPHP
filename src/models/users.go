package models

import (

)
import (
	"gopkg.in/mgo.v2/bson"
	"github.com/RhettDelFierro/GolangPHP/src/controllers"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserInfo struct{
	Id	bson.ObjectId	`bson:"_id, omitempty" json:"id"`
	UserName string		`json:"username"`
	Password string		`json:"password, omitempty"`
	Email	string		`json:"email"`
	HashPassword	[]byte	`json:"hashpassword, omitempty"`
}

func GetUser(user controllers.User) {
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return nil, err
	}
	defer session.Close()



}
//traight up take daata from json.
//adding a new user document into mongoDB.
func RegisterUser(user *UserInfo) error{

	obj_id := bson.NewObjectId()
	user.Id = obj_id

	hashpw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error generating hashpw")
		panic(err)
	}
	user.HashPassword = hashpw
	//so we don't store the unhashed pw
	user.Password = ""

	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return nil, err
	}
	defer session.Close()

	c := session.DB("taskdb").C("users")

	err = c.Insert(&user)
	return err
}