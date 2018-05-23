
package main

import (
	"fmt"
)

func someSubRoutine(threadNumber int) {
	for i := 0; i < 1000; i++ {
		fmt.Println(threadNumber, ":\t", i)
	}
}

func main() {
	go someSubRoutine(0)
	go someSubRoutine(1)

	var input string
	fmt.Scanln(&input)
}
