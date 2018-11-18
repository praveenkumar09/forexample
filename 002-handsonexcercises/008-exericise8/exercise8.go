package main

import "fmt"

const (
	a = 14
	b = 19.78
	c = "Jack hill"
)
const (
	d int     = 18
	e float64 = 20
	f string  = "Jack ma"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
