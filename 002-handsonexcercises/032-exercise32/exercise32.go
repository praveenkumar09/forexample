package main

import "fmt"

type person struct {
	first  string
	last   string
	flavor []string
}

func main() {
	p1 := person{
		first:  "Praveen",
		last:   "Kumar",
		flavor: []string{"Butter scotch", "Vannila", "Pistha"},
	}
	fmt.Println(p1)

	p2 := person{
		first:  "Manu",
		last:   "Prasad",
		flavor: []string{"Strawberry", "Blue berry", "ocean current"},
	}
	fmt.Println(p2)

	for i, v := range p1.flavor {
		fmt.Println(i, v)
	}

	for x, y := range p2.flavor {
		fmt.Println(x, y)
	}
}
