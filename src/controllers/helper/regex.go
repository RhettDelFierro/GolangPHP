package helper

import (
	//"reflect"
	"regexp"
	//"fmt"
	//"github.com/fatih/structs"
	//"github.com/RhettDelFierro/GolangPHP/src/controllers"
	//"strconv"
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"fmt"
	//"fmt"
)

func TestValidEntry(entry NewStudent) map[string]string{



	regex_tests := make(map[string]string)
	//32 alphanumeric characters. No spaces, but underscores allowed
	regex_tests["name"] = "/^[A-Za-z0-9_]{1,32}$/"
	regex_tests["course"] = "/^[A-Za-z0-9_]{1,32}$/"
	//only numbers
	regex_tests["grade"] = "/^100$|^[1-9]?[0-9]$/"
	regex_map := map[string]string{}

	//grade, _ := this.Grade(map[string]interface{})
	//this.Grade.(map[string]interface{})["value"]

	//gradeInt, _ := strconv.Atoi(this.Grade["value"])
	//grade := new(bytes.Buffer)
	//err := binary.Write(grade, binary.LittleEndian, gradeInt)
	//if err != nil {
	//	panic(err)
	//}

	//bs := make([]byte, grade)
	//binary.LittleEndian.PutUint32(bs, 31415926)
	//fmt.Println(bs)

	//DRY, but try and find a better solution.
	if boolean, _ := regexp.Match(regex_tests[entry.Name["description"]], []byte(entry.Name["value"])); boolean {
		regex_map["regex_name_error"] = entry.Name["invalid"]
	}
	//if boolean, _ := regexp.Match(regex_tests["grade"], []byte(this.Grade["value"])); boolean {
	//	regex_map["regex_course_error"] = this.Grade["invalid"]
	//}
	if boolean, _ := regexp.Match(regex_tests["grade"], []byte(entry.Name["value"])); boolean {
		regex_map["regex_course_error"] = entry.Grade["invalid"]
	}
	if boolean, _ := regexp.Match(regex_tests[entry.Course["description"]], []byte(entry.Course["value"])); boolean {
		regex_map["regex_grade_error"] = entry.Course["invalid"]
	}

	return regex_map
}