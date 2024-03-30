package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       int
	createAt  time.Time
}

// Method of the struct
func (user *User) printUserValue() {
	fmt.Println("------------------------------")
	fmt.Println("First Name : ", user.firstName)
	fmt.Println("Last Name : ", user.lastName)
	fmt.Println("Age : ", user.age)

	fmt.Println("------------------------------")
	fmt.Println()
}

func (user *User) clearUserName() {
	user.firstName = ""
}

func main() {

	appUser := User{
		firstName: "Virendra",
		lastName:  "patel",
		age:       24,
		createAt:  time.Now(),
	}

	appUserCreatedByAnotherWay := User{"Virendra", "Patel", 25, time.Now()}

	appUser.printUserValue()
	appUserCreatedByAnotherWay.printUserValue()

	appUserCreatedByAnotherWay.clearUserName()
	fmt.Println("After clear username")
	appUserCreatedByAnotherWay.printUserValue()

}
