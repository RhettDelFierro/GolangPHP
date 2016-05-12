package models

import(
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
	//"errors"
)

type ErrorString struct {
	E string
}

func (this ErrorString) Error() string {
	return this.E
}

func getDBConnection() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err !=nil{
		fmt.Println("error in DB connection")
		//panic(err)
		return nil, ErrorString{"there was an error in th DB connection"}
		//return nil, errors.New("error in DB connection")
	}
	return session, err
}