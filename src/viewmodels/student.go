package viewmodels

import (
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	Name   string	`json:"name"`
	Course string	`json:"course"`
	Grade  int	`json:"grade"`
	Id     bson.ObjectId `json:"id"`
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