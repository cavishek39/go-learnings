package main

import "fmt"

func swap (a string, b string) (string, string) {
    return b, a
}

func main() {
	a, b := swap("abc", "def")
	fmt.Println("Swapping", a, b)
}