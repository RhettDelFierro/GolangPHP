package helper

import (
	//"reflect"
	"regexp"
	//"fmt"
	//"fmt"
)

func TestValidEntry(entry map[string]map[string]string) []string {

	regex_tests := make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "^[A-Za-z0-9_]{1,32}$"
	regex_tests["course"] = "^[A-Za-z0-9_]{1,32}$"
	//only numbers
	regex_tests["grade"] = "^100$|^[1-9]?[0-9]$"
	var regex_array []string


	for key, value := range regex_tests{
		index := key
		//r,_ := regexp.Compile(value)
		for k, v := range entry {
			if k == index {
				//fmt.Println("here are k, index, and v[value]:", k, index, v["value"])
				if boolean, _ := regexp.MatchString(value, v["value"]); !(boolean) {
					//fmt.Println("boolean (it should be true: ", boolean)
					regex_array = append(regex_array, v["invalid"])
				}
			}
		}
	}

	return regex_array
}