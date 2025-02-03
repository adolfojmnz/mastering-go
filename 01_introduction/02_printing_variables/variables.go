package main

import (
	"fmt"
	"math"
)

var Global int = 7272
var AnotherGlobal = -5328

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value: ", j)
	fmt.Println("i value: ", i)
	j = Global
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("g:%d, i:%d, j=%d, k:%.2f.\n", Global, i, j, k)
}

