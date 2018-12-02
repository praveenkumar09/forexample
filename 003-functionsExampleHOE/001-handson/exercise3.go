package main

import "fmt"

func main() {
	defer thisRunsSecond()
	thisRunsFirst()
}

func thisRunsSecond() {
	fmt.Println("This gets called first but executes second")
}

func thisRunsFirst() {
	fmt.Println("This gets called second but executes first")
}
