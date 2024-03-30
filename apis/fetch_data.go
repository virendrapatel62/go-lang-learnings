package apis

import "fmt"

// to use a function outside of the package need to start with caps.
// do not use fetchData; use FetchData

func FetchData() {
	fmt.Println("Fetching data from anothor package")
}
