package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string // uppercase makes the properties non private and accessible
	LastName  string
	Score     int
}

// Create a Method on Instructor
// technically though not best practice this can in a different file as long as its the same package

func NewInstructor(name string, lastname string) Instructor {
	return Instructor{FirstName: name, LastName: lastname}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.LastName, i.FirstName, i.Score)
}
