package helper

import (
	"regexp"
)

func TestValidEntry(entry struct{}) []string {
	regex_tests := make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "/^[A-Za-z0-9_]{1,32}$/"
	regex_tests["course"] = "/^[A-Za-z0-9_]{1,32}$/"
	//only numbers
	regex_tests["grade"] = "/^100$|^[1-9]?[0-9]$/"

	var regex_array []string
	for name,_ := range entry {
		boolean, err := regexp.Match(regex_tests[name]["description"], name["value"])
		if err != nil{
			panic(err)
		}
		if !boolean {
			regex_array = append(regex_array, name)
		}
	}

	return regex_array
	//return preg_match($regex_tests[$key], $value);
}