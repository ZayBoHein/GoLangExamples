package main

import "fmt"

func main() {
	var name string

	fmt.Print("Please input your name:")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello there %s", name)
}
