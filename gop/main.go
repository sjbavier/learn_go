package main

import "fmt"

// can have multiple functions named init, they will be ran before main
func init() {
	fmt.Println("A")
}
func init() {
	fmt.Println("B")
}
func main() {
	var message string
	var price = 34.3

	defer fmt.Println("i'm defered in execution") // defered until end of function call
	defer fmt.Println("Bye") // defer is put on call stack
	defer fmt.Println("Good") // last in first out, also if there is a panic, it will execute defers before closing

	fmt.Println("safe line printing for all oses")
	print("running go", message, price)
	stateTax, someTax := CalculateTax(100.3)
	fmt.Println("some tax stuff", stateTax, someTax)
	age := 100
	 willNotIncrementAge(age)
	 willNotIncrementAge(age)
	 willIncrementAge(&age) // passing in a pointer or reference to variable age
	fmt.Println("called twice and copied twice", age)
}
