package helper
import (
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	"github.com/RhettDelFierro/GolangPHP/src/models"
)

func StudentsToViewModel(category models.Student) viewmodels.Student {
	result := viewmodels.Student{
		Name: category.Name(),
		Course: category.Course(),
		Grade: category.Grade(),
		Id: category.Id(),
	}

	return result
}