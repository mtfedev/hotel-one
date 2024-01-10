package main

import "fmt"

type User struct {
	Username string
	Age      int
}

func getNumber() int {
	return 88
}

func main() {
	user := User{
		Username: "Mike",
		Age:      getNumber(),
	}

	fmt.Println("The user is ", user)
}
