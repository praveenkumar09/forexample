package main

import "fmt"

func main() {
	x := 10
	y := 21
	if x < y {
		fmt.Println("X is less than y")
	}
	if y < x {
		fmt.Println("Y is less than X")
	}
	if y <= x {
		fmt.Println("Y is less than or equal to x")
	}
	if x <= y {
		fmt.Println("X is less than or equal to y")
	}
	if x != y {
		fmt.Println("X is not equal to Y")
	}
	if x == y {
		fmt.Println("X is equal to y")
	}
}
