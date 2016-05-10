package helper

import ()
import (
	"html"
	"strconv"
)

//the point of this is just a place to set error messages pretty much.
type NewStudent struct {
	Name       map[string]string
	Course     map[string]string
	Grade      map[string]interface{}
	Duplicate  map[string]interface{}
	Auth_token map[string]string
}

type StudentStruct struct {
	Name string
	Course string
	Grade int
}

func (this *StudentStruct) Convert(value map[string]string) {
	var err error
	this.Name = html.EscapeString(value["name"])
	this.Course = html.EscapeString((value["course"]))
	grade := html.EscapeString((value["grade"]))
	this.Grade, err = strconv.Atoi(grade)
	if err != nil {
		panic(err)
	}
}

//should I just unmarshall from json.Unmarshall in the controller to a variable of type NewStudent? No because of the types in the field of NewStudent.
func (this *NewStudent) Make(name string,course string, grade string){ //pass in a struct instead?
	this.Name["value"] = name
	this.Course["value"] = course
	this.Grade["value"] = grade
}

func (this *NewStudent) ErrorMaker(postData map[string]string) {

	var safeEntry *StudentStruct
	safeEntry.Convert(postData)

	this.Name = map[string]string{
		"value": safeEntry.Name,
		"description": "name",
		"invalid": "Invalid name, please use only letters and numbers",
		"error": "There was an error adding student to database.",
	}
	this.Course = map[string]string{
		"value": safeEntry.Course,
		"description": "course",
		"invalid": "Only numbers and letters please.",
		"error": "There was an error adding student to database",
	}
	this.Grade = map[string]interface{}{
		"value": safeEntry.Grade,
		"description": "grade",
		"invalid": "Only numbers 0-100",
		"error": "There was an error adding student to database",
	}
	this.Duplicate = map[string]interface{}{
		"value": "",//struct{},
		"description": "duplicate",
		"error": "Records show you've alread recorded this entry", //continue to add (if they're the same user that added the name in the first place)?
	}
	this.Auth_token = map[string]string{
		"value":	"",
		"description": "auth_token",
		"access_error":	"you do not have the privileges to add a student",
		"expired":	"Login session expired",
	}

}
