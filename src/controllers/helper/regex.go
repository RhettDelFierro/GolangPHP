package helper

import (
	"regexp"
)

func TestValidEntry(entry struct{}) bool {
	regex_tests = make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "/^[A-Za-z0-9_]{1,32}$/"
	regex_tests["course"] = "/^[A-Za-z0-9_]{1,32}$/"
	//only numbers
	regex_tests["grade"] = "/^100$|^[1-9]?[0-9]$/"

	for entry,field := range entry {
		if(regexp.Match(regex_tests[entry]))
	}

	return true
	//return preg_match($regex_tests[$key], $value);
}