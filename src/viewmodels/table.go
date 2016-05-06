package viewmodels

import ()

type Students struct {
	Name   string
	Course string
	Grade  int
	Id     int
}

type School struct {
	All []Students
}

func GetGrades(id int) Students {
	var student Students

	student1 := Students{Name: "Student1", Course: "Course1", Grade: 100, Id: 1,}
	student2 := Students{Name: "Student2", Course: "Course2", Grade: 100, Id: 2,}
	student3 := Students{Name: "Student3", Course: "Course3", Grade: 100, Id: 3,}

	switch id {
		case 1:
			student = student1

		case 2:
			student = student2

		case 3:
			student = student3
	}

	//class := School{All: []Students{student1, student2, student3}}
	//fmt.Println(class)

	return student


	//class := School{All: []Students{student1, student2, student3}}

	//return class
	//return class.All[id-1]
	//return class.All[id-1]
	//figure out a way to return the whole school when they want to populate it (maybe make that an option on sgt html)
}