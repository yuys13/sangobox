package main

import (
	"fmt"
)

func main() {
	message := getMessage()
	fmt.Println(message)
}

func getMessage() string {
	return "Hello, World!"
}
