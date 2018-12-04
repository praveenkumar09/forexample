package main

import "fmt"

type person struct {
	age int
}

func main() {
	p := person{20}
	changeMe(&p)
}

func changeMe(p *person) {
	p.age = 30
	fmt.Println((*p).age)
}
