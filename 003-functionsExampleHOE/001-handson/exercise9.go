package main

import "fmt"

func main() {
	ai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := even(sum, ai)
	fmt.Println(x)
}

func sum(ai []int) int {
	total := 0
	for _, v := range ai {
		total += v
	}
	return total
}

func even(f func(ai []int) int, vi []int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi)
}
