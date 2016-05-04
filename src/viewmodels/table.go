package viewmodels

import(

)

type Students struct{
	Name 	string
	Course 	string
	Grade 	int
}

type School struct{
	All []Students
}

func GetGrades() School {
	student1 := Students{Name: "Student1", Course: "Course1", Grade: 100,}
	student2 := Students{Name: "Student2", Course: "Course2", Grade: 100,}
	student3 := Students{Name: "Student1", Course: "Course3", Grade: 100,}
	class := School{All: []Students{student1, student2, student3}}

	return class
}