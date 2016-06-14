package models

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)
//students kept private
type Student struct {
	name   string	`bson:"name,omitempty"`
	course string	`bson:"course,omitempty"`
	grade  int 	`bson:"grade,omitempty"`
	id     bson.ObjectId `bson:"_id,omitempty"`
}
//type in DB.
type DBStudent struct {
	Name   string	`bson:"name,omitempty" json:"name"`
	Course string	`bson:"course,omitempty" json:"course"`
	Grade  int	`bson:"grade,omitempty" json:"grade"`
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id"`
}

type DeleteError struct {
	E string
}

func (this DeleteError) Error() string {
	return this.E
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
		//panic(err)
		return nil, err
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")
	school := []Student{}
	iter := c.Find(nil).Iter()
		result := DBStudent{}
			for iter.Next(&result) {
				var student Student
				student.SetName(result.Name)
				student.SetCourse(result.Course)
				student.SetGrade(result.Grade)
				student.SetId(result.Id)
				school = append(school, student)
			}
	if err = iter.Close(); err != nil { panic(err) }


	return school, err
}


func AddStudents(student *Student) ([]Student, error) {

	//converting the models.Student(fields not exposed) into DBStudent(exposed fields) The DB won't see the Student{} fields.
	dbStudent := DBStudent{student.Name(), student.Course(), student.Grade(), student.Id(),}

	session, err := getDBConnection()

	if err != nil {
		fmt.Println("Error in DB connection")
		//errors = append(errors, err) //just append the string "Error in DB connection?"
		//panic(err)
		return nil, err
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")

	//check for duplicate.
	//can use go routine and channels here possibly to send info.
	duplicates := duplicate(&dbStudent)
	if len(duplicates) != 0 {
		fmt.Println("From inside AddStduents, the duplicates):", duplicates)
		return duplicates, err
	} else {

		err = c.Insert(&dbStudent)
		if err != nil {
			fmt.Println("error in inserting")
			//panic(err)
			return nil, err
		}
	}

	return duplicates, err
}

//can throw an error if the id doesn't match anything. They may have possibly deleted it before
//and it may still be there in the DOM.
func DeleteStudents(id string) error {
	session, err := getDBConnection()

	if err != nil {
		//fmt.Println("Error in AddStudents function")
		//panic(err)
		return err
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")

	if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		//fmt.Println(err);
		return DeleteError{"There was an error deleting the student. This student may not exist in the DB. Please refresh the table."}
	} else {
		return nil
	}
}

func duplicate(student *DBStudent) []Student {
	session, err := getDBConnection()

	if err != nil {
		fmt.Println("Error in DB connection")
		//errors = append(errors, err) //just append the string "Error in DB connection?"
		panic(err)
	}
	defer session.Close()

	c := session.DB("taskdb").C("categories")
	school := []Student{}

	iter := c.Find(bson.M{"name": student.Name, "course": student.Course, "grade": student.Grade}).Iter()
		result := DBStudent{}
			for iter.Next(&result) {
				var student Student
				student.SetName(result.Name)
				student.SetCourse(result.Course)
				student.SetGrade(result.Grade)
				student.SetId(result.Id)
				school = append(school, student)
			}
	if err != nil {
		panic(err)
	}

	fmt.Println("from the duplicate function: ", school)

	return school
}

//I think this is more for a member
//func CreateSession(student Student) (Session, error) {
//	result := Session{}
//	result.Id = student.Id
//}