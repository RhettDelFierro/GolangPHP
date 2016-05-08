package models

import (
	//"crypto/sha256" //use oauth2
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)
//students kept private
type Student struct {
	name   string	`bson:"name,omitempty"`
	course string	`bson:"course,omitempty"`
	grade  int `bson:"grade,omitempty"`
	id     bson.ObjectId `bson:"_id,omitempty"`
}
//type in DB.
type DBStudent struct {
	Name   string	`bson:"name,omitempty"`
	Course string	`bson:"course,omitempty"`
	Grade  int	`bson:"grade,omitempty"`
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id"`
}

//getter
func (this *Student) Name() string { return this.name }
func (this *Student) Course() string { return this.course }
func (this *Student) Grade() int { return this.grade }
func (this *Student) Id() bson.ObjectId { return this.id }
//setters
func (this *Student) SetName(name string) { this.name = name }
func (this *Student) SetCourse(course string) { this.course = course }
func (this *Student) SetGrade(grade int) { this.grade = grade }
func (this *Student) SetId(id bson.ObjectId) { this.id = id }

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
		result := DBStudent{}
			for iter.Next(&result) {
				var student Student
				fmt.Printf(result.Name)
				student.SetName(result.Name)
				student.SetCourse(result.Course)
				student.SetGrade(result.Grade)
				student.SetId(result.Id)
				school = append(school, student)
			}
	if err = iter.Close(); err != nil { panic(err) }

	return school, err
}


func AddStudents(student *Student) error {
	dbStudent := DBStudent{student.Name(), student.Course(), student.Grade(), student.Id(),}

	session, err := getDBConnection()

	if err != nil {
		fmt.Println("Error in AddStudents function")
		panic(err)
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")

	err = c.Insert(&dbStudent)
	if err != nil {
		fmt.Println("error in inserting")
		panic(err)
	}

	return err
}

func DeleteStudents(id string) bool {
	session, err := getDBConnection()

	if err != nil {
		fmt.Println("Error in AddStudents function")
		panic(err)
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")

	if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		fmt.Println(err);
		return false
	} else {
		return true
	}
}

//I think this is more for a member
//func CreateSession(student Student) (Session, error) {
//	result := Session{}
//	result.Id = student.Id
//}