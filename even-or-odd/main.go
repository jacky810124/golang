package main

import (
	"fmt"
)

func main() {
	var numbers = []int{}

	for i := 0; i <= 10; i++ {
		// just want to create 0-10 integer slice automatically, not manually.
		numbers = append(numbers, i)
		print(i)
	}
}

func print(num int) {
	if num%2 == 0 {
		fmt.Println(num, "even")
	} else {
		fmt.Println(num, "odd")
	}
}
