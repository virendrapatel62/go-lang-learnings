package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age : ", age)

	// Case 1:
	// adultYears := getAdultsYear(age)
	// fmt.Println("Adult Years :", adultYears)

	// Case 2:
	fmt.Println("Age pointer: ", agePointer)
	fmt.Println("Age pointer value: ", *agePointer)
	fmt.Println("Adult value: ", getAdultsYearWithPointer(agePointer))
	fmt.Println("Orignal age value : ", age)

	// Case 3:

	getAdultsYearWithPointerManipulation(&age)

	fmt.Println("Orignal age value : ", age)

}

// Case 1:
// copy of int will be here
func getAdultsYear(age int) int {
	return age - 18
}

// Case 2
// no copy of int
func getAdultsYearWithPointer(age *int) int {
	return *age - 18
}

// Case 3
// no copy of int
// manipulating value using pointer
func getAdultsYearWithPointerManipulation(age *int) {
	*age = *age - 18
}
