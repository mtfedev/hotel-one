package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 unbuffered channel 0 book
	// 2 buffered channel 10 books
	//resultch := make(chan string, 10)  ->buffered channel
	resultch := make(chan string) // ->unbuffered channel
	resultch <- "foo"             // -> if now Full ->ITTT WILL BLOCK -> block HERE
	// this coe below will never execute !!!!
	result := <-resultch
	fmt.Println(result)

	//result := fetchResource(1)
	//fmt.Println("The result:", result)
	// asynk
	//go fetchResource()

}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
