package main

import "fmt"

func main() {

	fmt.Println(next(1))
}

// next returns the next number
func next(x int) int {
	return x + 1
}
