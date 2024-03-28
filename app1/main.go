package main

import "fmt"

func printNewLine() {
	fmt.Println()
}

const ADD_TODO = 1
const MARK_DONE = 2

func printOptions() {
	fmt.Println("Select an option please 👇")
	fmt.Println("👉 1. To add todo")
	fmt.Println("👉 2. To Mark a todo done")
}

func takeUserInput() int {
	var optionInput int = 0
	fmt.Scan(&optionInput)
	return optionInput
}

func performTask(selectedOption int) {
	switch selectedOption {
	case ADD_TODO:
		{

		}

	case MARK_DONE:
		{

		}

	}
}

func main() {
	fmt.Println("Welcome to my first TODO app in GOLANG 🚀 ")
	printNewLine()

	printOptions()
	var optionsInput int = takeUserInput()
	performTask(optionsInput)

}
