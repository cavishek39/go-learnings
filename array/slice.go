package main

import "fmt"

func main() {
	sample_array := [6]int{1, 2, 3, 4, 5, 6}

	var sliced_array []int = sample_array[3:5]

	fmt.Println(sliced_array)
}
