package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/app/apis"
	"github.com/app/services"
)

func main() {
	sum()       // from same main package
	substract() // same main package

	services.NotificationService() // from another package services
	apis.FetchData()               // form another package

	//using third party package

	fmt.Println(randomdata.FullName(1))
	fmt.Println(randomdata.FullName(1))
	fmt.Println(randomdata.FullName(1))
	fmt.Println(randomdata.FullName(1))
	fmt.Println(randomdata.FullName(1))

}
