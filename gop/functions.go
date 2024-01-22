package main

import (
	"fmt"
	"gop/data"
)

func PublicFunc() { // TitleCase is public and exported to other packages
	fmt.Println("public function to package")
	fmt.Println(data.TestData)
}

func privateFunc() { // camelCase is private to the package
	fmt.Println("private function to package")
}

func add(a int, b int) int { // arguments are passed as a copy not as a reference or pointer
	return a + b
}

func addAndSubtract(a int, b int) (int, int){ // can return multiple values, must return both at same time
	return a + b, a - b
}

func CalculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func CalculateTaxWithName(price float32) (stateTax float32, cityTax float32) { // named returns automatically return
	stateTax = price * 0.09
	cityTax = price * 0.02
	return 
}

func willNotIncrementAge(age int) {
	age++
}

func willIncrementAge(age *int){ // passing a pointer or reference to the variable
	if(*age > 140){
		panic("this raises a panic") // this is not for error handling, this is for escalation and halting
	}
	*age++ // changing the reference to age
}
