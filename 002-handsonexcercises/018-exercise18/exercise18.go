package main

import "fmt"

func main() {
	x := 4
	y := 5
	if x < y {
		fmt.Println("X is less than y")
	} else if y < x {
		fmt.Println("Y is less than X")
	} else {
		fmt.Println("Both are equal")
	}
}
