package data

import "fmt"

var Countries [10]string // array of 10 items of type string, traditional array with fixed length and data type
var Slice []int          // slice is an array of variable size.. array under the hood
var Mapping map[int]bool // key of int and value of bool
var ConstructedArray = [2]int{100, 500}

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "USA"
	Slice := append(Slice, 2) // we can append slices, this is creating a new variable Slice :=

	wellKnownPorts := map[string]int{"http": 80, "https": 443} // constructing a map
	quantity := len(Countries)
	capacity := cap(Countries)
	fmt.Println("Countries saved", quantity, capacity, ConstructedArray, Slice, wellKnownPorts)
}
