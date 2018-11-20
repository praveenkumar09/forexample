package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("it is true")
	case false:
		fmt.Println("it is false")
	default:
		fmt.Println("this is default")
	}
}
