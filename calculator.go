package main

import (
	"github.com/app/apis"
	"github.com/app/services"
)

func main() {
	sum()       // from same main package
	substract() // same main package

	services.NotificationService() // from another package services
	apis.FetchData()               // form another package

}
