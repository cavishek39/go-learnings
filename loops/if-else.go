package main

import "fmt"

func main() {
	i := 0
	limit := 5
	for i < limit {
		if i % 2 == 0 {
			fmt.Println("i is zero")
		} else {
			fmt.Println("i is not zero")
		}
		i += 1
	}
}