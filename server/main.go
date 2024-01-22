package main

import (
	"fmt"
	"server/data"
)

func main() {

	// two ways to implement structure
	max := data.Instructor{Id: 1, LastName: "Test"}
	// set properties after creation if they are captilalized and Public not private
	max.FirstName = "Maximiliano"
	swiftWS := data.NewWorkshop("Swift with iOS", max)

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals"}
	fmt.Println(max.Print(), goCourse)

	var courses [2]data.Signable // implicit implementation
	courses[0] = goCourse
	courses[1] = swiftWS

	for i, course := range courses {
		fmt.Println(i, course)
	}
}
