package converters

import (
	"github.com/RhettDelFierro/GolangPHP/src/models"
	"github.com/RhettDelFierro/GolangPHP/src/viewmodels"
	//"gopkg.in/mgo.v2/bson"
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