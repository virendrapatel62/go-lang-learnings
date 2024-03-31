package main

import (
	"fmt"

	products "github.com/arrays/models"
)

func main() {
	numbers := []int{1, 2, 4, 5, 6}
	fmt.Println(numbers)

	products := []products.Product{
		products.New("Dell", 99002), products.New("Mac book pro", 899949),
	}

	fmt.Println(products)

	fmt.Println(products[1].ProductName)

	// slices

	fmt.Println(numbers[0])
	fmt.Println(numbers[1:4]) // index 1-3
	fmt.Println(numbers[3:])  // index 3-n
	fmt.Println(numbers[:4])  // index 0-3

	slice1 := numbers[3:] // doesn't copy the content , just a reference

	fmt.Println("slice 1: ", slice1)

	slice1[1] = 999 // changing slice changes the original data as well.

	fmt.Println("slice 1: ", slice1)

	fmt.Println("numbers: ", numbers)

}
