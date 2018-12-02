package main

import "fmt"

func main() {
	ai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	xi := fool(ai...)
	yi := barl(ai)
	fmt.Println(xi)
	fmt.Println(yi)
}

func fool(ai ...int) int {
	total := 0
	for _, v := range ai {
		total += v
	}
	return total
}

func barl(ai []int) int {
	total := 0
	for _, v := range ai {
		total += v
	}
	return total
}
