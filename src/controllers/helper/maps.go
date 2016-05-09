package helper

import ()

//the point of this is just a place to set error messages pretty much.
type NewStudent struct {
	Name       map[string]string
	Course     map[string]string
	Grade      interface{}
	Duplicate  interface{}
	Auth_token map[string]string
}

type dataHandler interface {

}

//no short variable declaration because it's not in a function.
var AddingStudent = NewStudent{
	Name: {
		"value": "",
		"invalid": "Invalid name, please use only letters and numbers",
		"error": "There was an error adding student to database.",
	},
	Course: {
		"value": "",
		"invalid": "Only numbers and letters please.",
		"error": "There was an error adding student to database",
	},
	Grade: {
		"value": "",
		"invalid": "Only numbers please",
		"error": "There was an error adding student to database",
	},
	Duplicate: {
		"value": "",
		"error": "Records show you've alread recorded this entry", //continue to add (if they're the same user that added the name in the first place)?
	},
	Auth_token: {
		"value":	"",
		"access_error":	"you do not have the privileges to add a student",
		"expired":	"Login session expired",
	},
}
