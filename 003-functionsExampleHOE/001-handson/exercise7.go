package main

import "fmt"

func main() {
	x := food()
	fmt.Println(x)
}

func food() string {
	return fmt.Sprint("The string is coming from function foo")
}
