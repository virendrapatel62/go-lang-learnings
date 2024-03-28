package main

import (
	"fmt"
	"os"
)

func printNewLine() {
	fmt.Println()
}

const ADD_TODO = 1
const MARK_DONE = 2
const PRINT_TODOS = 3
const EXIT = 9

var todos []string = []string{"Do somthing", "more work"}

func printOptions() {
	fmt.Println("Select an option please 👇")
	fmt.Println("👉 1. To add todo")
	fmt.Println("👉 2. To Mark a todo done")
	fmt.Println("👉 3. To print all todos")
	fmt.Println("👉 9. EXIT ❌")
}

func takeUserInput() any {
	var input any
	fmt.Scan(&input)
	return input
}

func performTask(selectedOption int) {
	switch selectedOption {
	case ADD_TODO:
		{
			fmt.Println("Todo title : (without any space please 😫) : ")
			var todoToAdd string
			fmt.Scan(&todoToAdd)
			todos = append(todos, todoToAdd)

			break
		}

	case MARK_DONE:
		{
			fmt.Println("Enter the sequence number of todo to mark done: ")
			var serialNumber int
			fmt.Scan(&serialNumber)
			var index = serialNumber - 1
			todos = append(todos[:index], todos[serialNumber:]...)
			fmt.Println("Done ✅")
			break
		}

	case PRINT_TODOS:
		{
			printNewLine()

			fmt.Println("All the todos 🙌")

			for index, value := range todos {
				fmt.Printf("👉 %d. %s \n", index+1, value)
			}
			printNewLine()
			break
		}

	case EXIT:
		{
			os.Exit(0)
		}

	}
}

func main() {
	fmt.Println("Welcome to my first TODO app in GOLANG 🚀 ")
	printNewLine()

	for true {
		printOptions()
		var option = 0
		fmt.Scan(&option)
		performTask((option))
	}

}
