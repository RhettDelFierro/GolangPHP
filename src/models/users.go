package models

import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserInfo struct {
	Id           bson.ObjectId        `bson:"_id,omitempty" json:"id"`
	UserName     string                `json:"username"`
	Password     string                `json:"password,omitempty"`
	Email        string                `json:"email"`
	HashPassword []byte        `json:"hashpassword,omitempty"`
}

type RegisterError struct {
	E string
}

func (this RegisterError) Error() string {
	return this.E
}

func DuplicateUser(user UserInfo) (err error) {
	//just should make this whole getDBConnection and error handling block a reusable function.
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return err
	}

	defer session.Close()

	c := session.DB("taskdb").C("users")

	//if it finds one, it will write to u.
	u := UserInfo{}
	err = c.Find(bson.M{"username": user.UserName}).One(&u)
	if err != nil {
		fmt.Println("err in Duplicate: Find", err.Error())
		return err
	}

	return
}

//straight up take data from json.
//adding a new user document into mongoDB.
func RegisterUser(user *UserInfo) error {

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

	//check duplicate
	//err = c.Find(bson.M{"email": user.Email}).One(&user)

	//if err is of type not found, go on and register the user. But if it isn't then throw in the panic.
	if err != nil {
		panic(RegisterError{"Records show there is already a user with this email address. Please use another."})
	} else {
		err = c.Insert(user)
		if err != nil {
			fmt.Println("error registering: ", err)
			panic(err)
		}
	}

	return err
}

func LoginUser(user UserInfo) (u UserInfo, err error) {
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return u, err
	}
	defer session.Close()
	c := session.DB("taskdb").C("users")
	//want to check both the username and password. If either are successful, proceed with hashpassword process.
	_, err1 := CheckUser(user)
	_, err2 := CheckEmail(user)

	switch {
	case err1 == nil:
		err = c.Find(bson.M{"username": user.UserName}).One(&u)
		if err != nil {
			panic(err);
		}
	case err2 == nil:
		err = c.Find(bson.M{"email": user.Email}).One(&u)
		if err != nil {
			panic(err);
		}
	}

	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = UserInfo{}
		fmt.Println("Hashpassword error")
	}

	return
}

//for logging in.
func CheckEmail(user UserInfo) (u UserInfo, err error) {
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return u, err
	}
	defer session.Close()

	c := session.DB("taskdb").C("users")

	err = c.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			u = UserInfo{}
			return
		} else {
			panic(err);
		}
	}

	return

}

func CheckUser(user UserInfo) (u UserInfo, err error) {
	session, err := getDBConnection()

	if err != nil {
		//panic(err)
		return u, err
	}
	defer session.Close()

	c := session.DB("taskdb").C("users")
	err = c.Find(bson.M{"username": user.UserName}).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			u = UserInfo{}
			return
		} else {
			panic(err);
		}
	}

	return u, err

}