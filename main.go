package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello word")

	fmt.Println(add(5, 3))

	fmt.Println(strings.ToLower("MAM fajne CiUChY"))
}

func add(a int, b int) int {
	return a + b
}
