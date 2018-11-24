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

	mp := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(mp["Kumar"])

	for i, v := range p1.flavor {
		fmt.Println(i, v)
	}

	for x, y := range p2.flavor {
		fmt.Println(x, y)
	}
	fmt.Println("About to print from the map")
	for _,b := range mp{
		fmt.Println(b.first)
		for c,d := range b.flavor{
			fmt.Println(c,d)
			fmt.Println("----")
		}
	}
}
