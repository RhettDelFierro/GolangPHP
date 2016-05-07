package models

import (
	//"crypto/sha256" //use oauth2
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

//more for logged in members
//type Session struct {
//	Id int
//	Name string
//	Description string
//}

type Student struct {
	name   string	`bson:"name,omitempty"`
	course string	`bson:"course,omitempty"`
	grade  int `bson:"grade,omitempty"`
	id     bson.ObjectId `bson:"_id,omitempty"`
}

//type School struct {
//	Students []Student
//}


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

func (this *Student) Id() bson.ObjectId {
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

func (this *Student) SetId(id bson.ObjectId) {
	this.id = id
}

//func GetStudents() []Student { going to use the database instead.
//	student1 := Student{name: "Student1", course: "Course1", grade: 100, id: 1,}
//	student2 := Student{name: "Student2", course: "Course2", grade: 100, id: 2,}
//	student3 := Student{name: "Student3", course: "Course3", grade: 100, id: 3,}
//
//	class := []Student{student1, student2, student3}
//
//	return class
//}

func GetStudents() ([]Student, error) {
	session, err := getDBConnection()

	if err != nil {
		panic(err)
		//return Student{}, errors.New("Unable to connect to DB")
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")
	school := []Student{}
	iter := c.Find(nil).Iter()
		result := Student{}
			for iter.Next(&result) {
				fmt.Printf(result.name)
				//result.SetName(result.name)
				//result.SetCourse(result.course)
				//result.SetGrade(result.grade)
				//result.SetId(result.id)
				school = append(school, result)
			}
	if err = iter.Close(); err != nil {
		panic(err)
	}


	//fmt.Println("from the database get grades:")

	return school, err

}


func AddStudents(student *Student) {
	session, err := getDBConnection()

	if err != nil {
		fmt.Println("Error in AddStudents function")
		panic(err)
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")

	//student.SetId(bson.NewObjectId())

	fmt.Println("before student is added:", student)

	//student1 := Student1{"blah", "blah blah", 1, bson.NewObjectId()}

	err = c.Insert(student)
	if err != nil {
		fmt.Println("error in inserting")
		panic(err)
	}

}

//I think this is more for a member
//func CreateSession(student Student) (Session, error) {
//	result := Session{}
//	result.Id = student.Id
//}