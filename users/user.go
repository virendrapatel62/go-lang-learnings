package users

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` // starting with Caps mean it is accessible in another package
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	CreateAt  time.Time `json:"created_at"`
}

func (user *User) PrintUserValue() {
	fmt.Println("------------------------------")
	fmt.Println("First Name : ", user.FirstName)
	fmt.Println("Last Name : ", user.LastName)
	fmt.Println("Age : ", user.Age)

	fmt.Println("------------------------------")
	fmt.Println()
}

func (user *User) ClearUserName() {
	user.FirstName = ""
}

func New(firstName, lastName string, age int) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		CreateAt:  time.Now(),
		Age:       age,
	}
}

func (user User) SaveToFile() {
	fileName :=
		strings.ToLower(strings.Join([]string{user.FirstName, user.LastName}, "_") + ".json")
	jsonData, _ := json.Marshal(user)
	os.WriteFile(fileName, jsonData, 0664)
}
