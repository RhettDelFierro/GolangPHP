package helper
import (
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/models"
)

type DataHandler interface {
	Make(name string,course string, grade int)
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

//func DataInterfaceMapFunction(h DataHandler, name string, course string, grade int) bool {
//	h.Make(name, course, grade) //is it returning a struct or handler?
//	return true
//}