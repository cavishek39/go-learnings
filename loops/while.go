package main

import "fmt"

func main() {
	sum := 0

	// Create a while loop

	for sum < 10 {
		sum += 1
	}

	fmt.Println("Sum using while loop", sum)
}