package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", y, y, y)
}
