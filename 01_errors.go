package main

import (
	"errors"
	"fmt"
)

func check(a, b int) error {
	if a == 0 || b == 0 {
		return errors.New("Values must be greater than zero")
	}
	return nil
}

func main() {
	err := check(0, 2)
	fmt.Print(err)
}

