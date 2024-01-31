package main

import "fmt"


func split(sum int) (x, y int) {
	x = sum + 2
	y = sum + 3
	return
}

func main () {
	x, y := split(10)
    fmt.Println(x, y)
}