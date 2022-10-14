package main

import "fmt"

func main() {
	fmt.Println(next(2))
}

// next returns the next number
func next(x int) int {
	return x + 1
}
