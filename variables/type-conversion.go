package main

import "fmt"

func main() {
	var i int = 60

	var f float64 = float64(i)

	var u uint = uint(f)

	fmt.Println(i, f, u)

	fmt.Printf("Type of u is %T\n", u)
}