package main

import (
	"strconv"
)

func main() {
	variableAndDataType()
}

func variableAndDataType() {
	var name string = "Prieyudha Akadita S"
	var age int = 23
	var isMale bool = true
	var hobbies [3]string = [3]string{"eating", "sleeping", "reading book"}
	var message = name + " is " + strconv.Itoa(age) + " years old, [is male] " + strconv.FormatBool(isMale)
	println(message)
	println(hobbies[0])
}
