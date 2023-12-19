package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User{
	ID: 1, 
	FirstName: "Jehuda", 
	LastName: "Rajasa", 
	Email: "jehudarajasa@gmail.com",
}

func updateEmail(u *User, newEmail string) {
	u.Email = newEmail
	fmt.Println("in update email:", u.Email)
}

func main() {
	updateEmail(&u, "jehuda@zero-one-group.com");
	fmt.Println("Updated user: ", u)
}