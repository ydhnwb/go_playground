package main

import (
	"fmt"
	"strconv"

	"github.com/ydhnwb/go_playground/arr"
)

func main() {
	sliceOrArray()
}

func variableAndDataType() {
	//we can use :=, but remember using := it means declare and assign
	// = it oonly assign
	//example:
	//name := "Prieyudha"
	//name = "Udin Gambut"
	var name string = "Prieyudha Akadita S"
	var age int = 23
	var isMale bool = true
	var hobbies [3]string = [3]string{"eating", "sleeping", "reading book"}
	var message = name + " is " + strconv.Itoa(age) + " years old, [is male] " + strconv.FormatBool(isMale)
	println(message)
	println(hobbies[0])
}

func decisions() {
	age := 20
	if age >= 20 {
		println("You are old enough")
	} else if age > 15 {
		println("Still teen")
	} else {
		println("You are a kid")
	}
}

func loop() {
	//go only have one type of loop: for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
}

func sliceOrArray() {
	myProducts := arr.Array()
	fmt.Printf("%v\n", myProducts)
	myProducts = arr.RemoveWithKeepOrder(myProducts, 1)
	println("After removed...")
	fmt.Printf("%v\n", myProducts)
}
