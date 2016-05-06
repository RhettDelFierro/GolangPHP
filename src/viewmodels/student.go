package viewmodels

import (

)

type Student struct {
	Name   string
	Course string
	Grade  int
	Id     int
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