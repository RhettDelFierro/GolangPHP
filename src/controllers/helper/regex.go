package helper

import (
	"reflect"
	"regexp"
	"fmt"
)

func TestValidEntry(entry interface{}) []string {

	//reflect &NewStudent type to be a map
	//then iterate over the reflection map.
	//reflect &NewStudent type to be a map
	//then iterate over the reflection map.
	s := reflect.ValueOf(entry)
	typ := reflect.TypeOf(entry)
	if typ.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	//fmt.Println(typ)
	//right now it's a struct being reflected.
	//iterate through it:
	regex_tests := make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "/^[A-Za-z0-9_]{1,32}$/"
	regex_tests["course"] = "/^[A-Za-z0-9_]{1,32}$/"
	//only numbers
	regex_tests["grade"] = "/^100$|^[1-9]?[0-9]$/"

	var regex_array []string

	for key, value := range regex_tests {
		regex,_ := regexp.Compile(value)
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if key == f["description"] {
				if !regex.Match([]byte(f["value"])) {
					regex_array = append(regex_array, f["invalid"])
				}
			}
		}
	}
	return regex_array
}