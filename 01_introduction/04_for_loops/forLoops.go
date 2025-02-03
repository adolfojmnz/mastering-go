package main

import (
	"fmt"
)

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	j := 0
	for {
		if j == 10 {
			break
		}
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	// Iterating over a slice with for and range
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "value: ", v)
	}
}

