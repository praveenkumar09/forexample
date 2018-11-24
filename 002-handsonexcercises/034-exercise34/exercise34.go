package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 3,
			color: "Green",
		}, fourWheel: false,
	}
	fmt.Println("Printing out the entire t1 object :", t1)
	fmt.Println("Print the t1 vehicle doors :", t1.doors)
	fmt.Println("Print the t1 vehicle color :", t1.color)
	fmt.Println("Print the t1 fourWheel:", t1.fourWheel)

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Gold",
		}, luxury: true,
	}

	fmt.Println("Printing out the entire s1 object :", s1)
	fmt.Println("Print the s1 vehicle doors :", s1.doors)
	fmt.Println("Print the s1 vehicle color :", s1.color)
	fmt.Println("Print the s1 fourWheel:", s1.luxury)
}
