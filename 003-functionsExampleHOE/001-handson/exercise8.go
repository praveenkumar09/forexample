package main

import "fmt"

func main() {
	x := funcToReturnAFunc()
	fmt.Println(x())
}

func funcToReturnAFunc() func() int {
	return func() int {
		return 45
	}
}
