package helper

import ()
import (
	"html"
	//"strconv"
	//"fmt"
)

//the point of this is just a place to set error messages pretty much.
type NewStudent struct {
	Name       map[string]string
	Course     map[string]string
	Grade      map[string]string
	Duplicate  map[string]string
	Auth_token map[string]string
}

type StudentStruct struct {
	Name string
	Course string
	Grade string
}

func (this *StudentStruct) Convert(value map[string]string) {
	this.Name = html.EscapeString(value["name"])
	this.Course = html.EscapeString((value["course"]))
	this.Grade = html.EscapeString((value["grade"]))
}

//should I just unmarshall from json.Unmarshall in the controller to a variable of type NewStudent? No because of the types in the field of NewStudent.
func (this *NewStudent) Make(name string,course string, grade string){ //pass in a struct instead?
	this.Name["value"] = name
	this.Course["value"] = course
	this.Grade["value"] = grade
}

func ErrorMaker(postData map[string]string) map[string]map[string]string{
	//fmt.Println("inside errormaker: ", postData)
	var safeHTML StudentStruct
	safeEntry := &safeHTML
	safeEntry.Convert(postData)

	errorMap := make(map[string]map[string]string)

	errorMap["name"] = map[string]string{
		"value": safeEntry.Name,
		"description": "name",
		"invalid": "Invalid name, please use only letters and numbers",
	}

	errorMap["course"] = map[string]string{
		"value": safeEntry.Course,
		"description": "name",
		"invalid": "Invalid course name, please use only letters and numbers",
	}

	errorMap["grade"] = map[string]string{
		"value": safeEntry.Grade,
		"description": "grade",
		"invalid": "Only numbers 0-100",
	}

	errorMap["duplicate"] = map[string]string{
		"value": "",//struct{},
		"description": "duplicate",
		"error": "Records show you've already recorded this entry", //continue to add (if they're the same user that added the name in the first place)?
	}

	errorMap["auth_token"] = map[string]string{
		"value":	"",
		"description": "auth_token",
		"access_error":	"you do not have the privileges to add a student",
		"expired":	"Login session expired",
	}

	return errorMap
}
