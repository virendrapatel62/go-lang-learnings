package main

import (
	"fmt"

	"github.com/app/users"
)

// Method of the struct

func main() {

	appUser := users.New("Virendra", "patel", 24)

	appUserCreatedByAnotherWay := users.New("Virendra", "Patel", 25)

	appUser.PrintUserValue()
	appUserCreatedByAnotherWay.PrintUserValue()

	fmt.Println(appUser.FirstName)

	appUserCreatedByAnotherWay.ClearUserName()
	fmt.Println("After clear username")
	appUserCreatedByAnotherWay.ClearUserName()

}
