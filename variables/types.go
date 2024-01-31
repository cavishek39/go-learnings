package main

import "fmt"

// We can define multiple variables and assign them
var (
	MaxInt uint64 = 1 << 64 - 1 // default value
	isBoss bool = false
)

func main() {

	fmt.Println(MaxInt, isBoss)
}	