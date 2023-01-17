package main

import (
	"fmt"
)

func Factorial(n int) (v int) {
	if n > 0 {
		v = n * Factorial(n-1)
		return v
	}
	return 1
}
func main() {
	var n int
	fmt.Scan(&n)
	result := Factorial(n)
	fmt.Println(result)
}
