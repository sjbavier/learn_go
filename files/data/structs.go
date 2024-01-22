package data

import "fmt"

type User struct {
	id int
	name string
}


// We can have Methods added to the struct, but no inheritance
func (u User) PrettyPrint() string {
	return string(u.id) + ": " + u.name
}


func init() {
	var u1 User
	u1 = User {id: 1, name: "Go Club"} // with named properties
	u2 := User {2, "Golang"} // without name
	// without passing in the properties it defaults to 0 for int, "" for string and nil for 
	fmt.Println(u1, u2)
	
	u3 := User {3, "test"}

	fmt.Println(u3.PrettyPrint())
}