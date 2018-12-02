package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Praveen",
		last:  "Kumar",
		age:   26,
	}
	spokenString := p1.speak(p1.first, p1.last, p1.age)
	fmt.Println(spokenString)
}

func (p person) speak(fn string, ln string, age int) string {
	return fmt.Sprint(fn, " ", ln, " is a guy with age :", age)
}
