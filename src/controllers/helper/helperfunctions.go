package helper
import (
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"html"
	"strconv"
)

type DataHandler interface {
	ErrorMaker(name string,course string, grade int, auth_token string)
}

func StudentsToViewModel(category models.Student) viewmodels.Student {
	result := viewmodels.Student{
		Name: category.Name(),
		Course: category.Course(),
		Grade: category.Grade(),
		Id: category.Id(),
	}

	return result
}

//just take a struct?
func Convert(value map[string]string) struct{}{
	type Converted struct{
		name string
		course string
		grade int
	}
	var convertedStruct Converted
	convertedStruct.name = html.EscapeString(value["name"])
	convertedStruct.course = html.EscapeString((value["course"]))
	grade := html.EscapeString((value["grade"]))
	convertedStruct.grade = strconv.Atoi(grade)

	return convertedStruct
}

//func DataInterfaceMapFunction(h DataHandler, name string, course string, grade int) bool {
//	h.Make(name, course, grade) //is it returning a struct or handler?
//	return true
//}

//func TakeCareOfPOST(){
//
//}