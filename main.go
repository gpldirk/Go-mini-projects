package main

import "fmt"

func main() {
	fmt.Printf("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s", name)
}
