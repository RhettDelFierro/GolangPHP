package helper

import (
	//"reflect"
	"regexp"
	//"fmt"
	//"github.com/fatih/structs"
	//"github.com/RhettDelFierro/GolangPHP/src/controllers"
)

func (this *NewStudent) TestValidEntry() map[string]string{

	regex_tests := make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "/^[A-Za-z0-9_]{1,32}$/"
	regex_tests["course"] = "/^[A-Za-z0-9_]{1,32}$/"
	//only numbers
	regex_tests["grade"] = "/^100$|^[1-9]?[0-9]$/"
	var regex_map = make(map[string]string)


	//DRY, but try and find a better solution.
	if boolean, _ := regexp.Match(regex_tests[this.Name["description"]], []byte(this.Name["value"])); boolean {
		regex_map["regex_name_error"] = this.Name["invalid"]
	}
	if boolean, _ := regexp.Match(regex_tests[this.Grade["description"]], []byte(this.Name["value"])); boolean {
		regex_map["regex_course_error"] = this.Name["invalid"]
	}
	if boolean, _ := regexp.Match(regex_tests[this.Course["description"]], []byte(this.Name["value"])); boolean {
		regex_map["regex_grade_error"] = this.Name["invalid"]
	}

	return regex_map
}