package users

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string // starting with Caps mean it is accessible in another package
	LastName  string
	Age       int
	CreateAt  time.Time
}

func (user *User) PrintUserValue() {
	fmt.Println("------------------------------")
	fmt.Println("First Name : ", user.FirstName)
	fmt.Println("Last Name : ", user.LastName)
	fmt.Println("Age : ", user.Age)

	fmt.Println("------------------------------")
	fmt.Println()
}

func (user *User) ClearUserName() {
	user.FirstName = ""
}

func New(firstName, lastName string, age int) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		CreateAt:  time.Now(),
		Age:       age,
	}
}
