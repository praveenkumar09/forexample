package main

import "fmt"

func main() {
	fmt.Println("Comeon baby i am in")
	foo()
	fmt.Println("Do something more here")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Im in foo")
}
