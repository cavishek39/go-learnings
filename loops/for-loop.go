package main

import "fmt"

func forContinued(init, limit int) int{
	sum := 0

	for ; init <= limit; {
		sum += init
        init++
	}

	return sum
} 

func main() {
	sum := 0

	for i:= 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println("Sum using normal for-loops", sum)

	fmt.Println("Sum using for continued loop", forContinued(0, 10))
}