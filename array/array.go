package main

import "fmt"

// The ways we can declare an array in Go
func main() {
	// 1st way
	var a [2]int
	a[0] = 1
	a[1] = 2

	fmt.Println(a)

	// 2nd way
	array := [2]string{"foo", "bar"}
	fmt.Println(array)
}
