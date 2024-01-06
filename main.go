package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	//otherNumbers := make([]int, 12)

	numbers = append(numbers, 1)
	numbers = append(numbers, 13)
	fmt.Println(numbers)
	//fmt.Println(otherNumbers)

	//	users := map[string]int{}
	//users := make(map[string]int)

	//users["foo"] = 10
	//users["bar"] = 12

	//for key := range users {
	//	fmt.Printf("the key %s\n", key)
	//}

	//age, ok := users["bar"]
	//if !ok {
	//	fmt.Println("baz not exist in the map")
	//} else {
	//	fmt.Println("exist in the map", age)}

}
