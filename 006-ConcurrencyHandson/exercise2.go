package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("The person here is :", p.first, p.last, "  and he is ready to speak")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Praveen",
		last:  "Kumar",
		age:   27,
	}

	saySomething(p1) // doesnt work
	saySomething(&p1)

}
