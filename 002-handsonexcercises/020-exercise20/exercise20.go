package main

import "fmt"

func main() {
	favSport := "Cricket"
	switch favSport {
	case "Cricket":
		fmt.Println("it is cricket")
	case "Football":
		fmt.Println("it is football")
	default:
		fmt.Println("this is default")
	}
}
