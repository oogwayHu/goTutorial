package main

import (
	"fmt"
)

func main() {
	fmt.Println("5 + 6=", mySum(5, 6))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
