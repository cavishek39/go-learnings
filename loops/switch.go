package main

import (
	"fmt" 
 	"runtime"
)

func main() {
	switch v := runtime.GOOS; v {
		case "darwin":
			fmt.Println("macOS")
        case "linux":
			fmt.Println("Linux")
		default:
			fmt.Println("Windows")
	}
}