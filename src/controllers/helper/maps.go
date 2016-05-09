package helper

import ()

//the point of this is just a place to set error messages pretty much.
type NewStudent struct {
	Name       map[string]string
	Course     map[string]string
	Grade      map[string]interface{}
	Duplicate  map[string]interface{}
	Auth_token map[string]string
}

//should I just unmarshall from json.Unmarshall in the controller to a variable of type NewStudent? No because of the types in the field of NewStudent.
func (this *NewStudent) Make(name string,course string, grade string){ //pass in a struct instead?
	this.Name["value"] = name
	this.Course["value"] = course
	this.Grade["value"] = grade
}

//no short variable declaration because it's not in a function.
var AddingStudent NewStudent = NewStudent{
	Name: map[string]string{
		"value": "",
		"invalid": "Invalid name, please use only letters and numbers",
		"error": "There was an error adding student to database.",
	},
	Course: map[string]string{
		"value": "",
		"invalid": "Only numbers and letters please.",
		"error": "There was an error adding student to database",
	},
	Grade: map[string]interface{}{
		"value": "",
		"invalid": "Only numbers please",
		"error": "There was an error adding student to database",
	},
	Duplicate: map[string]interface{}{
		"value": "",
		"error": "Records show you've alread recorded this entry", //continue to add (if they're the same user that added the name in the first place)?
	},
	Auth_token: map[string]string{
		"value":	"",
		"access_error":	"you do not have the privileges to add a student",
		"expired":	"Login session expired",
	},
}