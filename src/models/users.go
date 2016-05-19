package models

import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserInfo struct{
	Id	bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	UserName string		`json:"username"`
	Password string		`json:"password,omitempty"`
	Email	string		`json:"email"`
	HashPassword	[]byte	`json:"hashpassword,omitempty"`
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
		return err
	}
	defer session.Close()

	c := session.DB("taskdb").C("users")
	fmt.Println(user.UserName)
	err = c.Insert(user)
	return err
}

func CheckUser(user *UserInfo)  (u UserInfo, err error){
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return nil, err
	}
	defer session.Close()
	c := session.DB("taskdb").C("users")
	err = c.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		fmt.Println("no records")
		return
	}
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = UserInfo{}
	}

	return

}