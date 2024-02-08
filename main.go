package main

import (
	"fmt"
	"time"
)

func main() {
	textch := make(chan string, 128)
	textch <- "A"
	textch <- "G"
	textch <- "F"
	close(textch)

	//	for {
	//		msg, ok := <-textch
	//		if !ok {
	//			break
	//		}
	//	}
	for msg := range textch {
		fmt.Println("the message =", msg)
	}
}

func fetchResourse(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprint("result %d", n)
}
