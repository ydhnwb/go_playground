package models

// Student is a data class
type Student struct {
	id     int
	name   string
	age    int
	isMale bool
}

// Create is a method to generate a new instance of student
func Create(id int, name string, age int, isMale bool) Student {
	aStudent := Student{id, name, age, isMale}
	return aStudent

}
