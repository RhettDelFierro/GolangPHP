package models

import(
	"gopkg.in/mgo.v2"
	"fmt"
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
		return nil, ErrorString{"there was an error in the DB connection"}
		//return nil, errors.New("error in DB connection")
	}
	return session, err
}