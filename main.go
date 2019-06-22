package main

import (
	"fmt"
	"time"
)

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d  pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	fmt.Println(time.Now())
}
