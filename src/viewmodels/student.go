package viewmodels

import (
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	Name   string
	Course string
	Grade  int
	Id     bson.ObjectId
}

type School struct {
	Students []Student
}

func GetStudents() School {
	result := School{}
	return result
}
/*
func AddStudent() Student {
	result := Student{}

	return result
}
*/