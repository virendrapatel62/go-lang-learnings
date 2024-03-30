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

func main() {

	appUser := User{
		firstName: "Virendra",
		lastName:  "patel",
		age:       24,
		createAt:  time.Now(),
	}

	appUserCreatedByAnotherWay := User{"Virendra", "Patel", 25, time.Now()}

	printUserValue(appUser)
	printUserValue(appUserCreatedByAnotherWay)

}

func printUserValue(user User) {
	fmt.Println("------------------------------")
	fmt.Println("First Name : ", user.firstName)
	fmt.Println("Last Name : ", user.lastName)
	fmt.Println("Age : ", user.age)

	fmt.Println("------------------------------")

	fmt.Println()

}
