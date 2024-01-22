package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator extreme")
	fmt.Println("******************")
	fmt.Println("Enter operation")
	fmt.Scanf("%s", &operation) // modify by sending address of variable
	fmt.Println("Enter number1")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter number2")
	fmt.Scanf("%d", &number2)
	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}
