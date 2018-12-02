package main

import "fmt"

func main() {
	func(s int) {
		for ; s > 0; s-- {
			fmt.Println(s)
		}
	}(100)
}
