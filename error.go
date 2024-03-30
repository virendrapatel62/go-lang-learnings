package main

import "fmt"

const throwError = true

func main() {
	panic("Hy there is an error ")
	fmt.Println("After Panic")
}
