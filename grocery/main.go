package main

import "fmt"

func main() {
	fmt.Println("hello, what is your name.")
	var name int
	fmt.Scanln(&name)
	fmt.Printf("Hi, %v!", name)
}