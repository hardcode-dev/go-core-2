package main

import "fmt"

const greeting = "Hello, World!"

func main() {
	printMsg(greeting)
}

func printMsg(msg string) {
	fmt.Println(msg) // плохо, а почему?
}
