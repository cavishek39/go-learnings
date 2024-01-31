package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
    fmt.Println("Sum => %d", add(1, 2))
	fmt.Println("Sub => %d", sub(1, 2))
}