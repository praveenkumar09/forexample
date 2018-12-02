package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func foo() int {
	fmt.Println("In foo method")

	return 42
}

func bar() (int, string) {
	return 420, fmt.Sprintln("Come on buddy im coming from func bar")
}
