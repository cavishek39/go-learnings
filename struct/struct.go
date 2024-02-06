package main

import "fmt"

type Error struct {
	message string
	statusCode int
}

func (e *Error) Error() string {
    return e.message
}

func main() {
	err := &Error{
		message: "Something went wrong",
        statusCode: 500,
	}

	fmt.Println(err)
}