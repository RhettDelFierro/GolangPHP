package models

import (

)

type Student struct {
	name   string
	course string
	grade  int
	id     int
}

type School struct {
	Students []Student
}


var School1 School

//getter
func (this *Student) Name() string {
	return this.name
}

func (this *Student) Course() string {
	return this.course
}

func (this *Student) Grade() int {
	return this.grade
}

func (this *Student) Id() int {
	return this.id
}

func (this *Student) SetName(name string) {
	this.name = name
}

func (this *Student) SetCourse(course string) {
	this.course = course
}

func (this *Student) SetGrade(grade int) {
	this.grade = grade
}

func (this *Student) SetId(id int) {
	this.id = id
}

func GetStudents() []Student {
	student1 := Student{name: "Student1", course: "Course1", grade: 100, id: 1,}
	student2 := Student{name: "Student2", course: "Course2", grade: 100, id: 2,}
	student3 := Student{name: "Student3", course: "Course3", grade: 100, id: 3,}

	class := []Student{student1, student2, student3}

	return class
}

func (this *School) AddStudents(student Student) {
	this.Students = append(student)
}