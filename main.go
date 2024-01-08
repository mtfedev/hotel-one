package main

import (
	"fmt"
)

func main() {
	name := "Alice"

	switch name {
	case "Alice":
		fmt.Println("The name Alice")
	default:
		fmt.Println("The name is default =>", name)
	}

	//users := map[string]int{
	//	"foo":  1,
	//	"dex":  2,
	//	"baf":  3,
	//	"max":  44,
	//	"zkix": 73,
	//}
	//for key, value := range users {
	//	fmt.Println("key %s value %d\n", key, value)
	//}

	//names := []string{"a", "b", "c", "d"}
	//for _, name := range names {
	//	if name == "a" {
	//		break
	//	}

	//}
	//fmt.Println("break out of loop")

	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])

}
