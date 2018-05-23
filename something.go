
package main

import (
	"fmt"
	"strconv"
)

func someSubRoutine(messages chan <- string, threadNumber int) {
	for i := 0; i < 100; i++ {
		go func() {
			s := strconv.Itoa(threadNumber) + ": " + strconv.Itoa(i)
			fmt.Println("message in ", s)
			messages <- s}()
	}
}

func main() {
	messages := make(chan string)
	go someSubRoutine(messages, 0)
	go someSubRoutine(messages, 1)

	for i :=0; i < 100; i++ {
		msg := <-messages

		fmt.Println("message out", msg)
	}
}
