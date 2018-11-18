package main

import "fmt"

const (
	a = 2014 + iota
	b = a + iota
	c = a + iota
	d = a + iota
	e = a + iota
)

func main() {
	fmt.Printf("%d\t\t%d\t\t%d\t\t%d\t\t", b, c, d, e)
}
