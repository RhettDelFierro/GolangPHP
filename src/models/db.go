package models

import(
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
)



func getDBConnection() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err !=nil{
		fmt.Println("error in DB connection")
		panic(err)
	}
	return session, err
}