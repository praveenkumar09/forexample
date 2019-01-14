package main

import "fmt"

func main() {
	go foo1()
	bar1()
	fmt.Println("Code about to exit")

}

func foo1() {
	fmt.Println("Hai there")
}

func bar1() {
	fmt.Println("Hai secnd there")
}
