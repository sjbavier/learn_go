package data

import "fmt"

type Duration float32

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) String() string { // custom to string method, override?
	return fmt.Sprintf("---- %v ---- %v", c.Name, c.Instructor.FirstName)
}

func (c Course) SignUp() bool {
	return true
}
