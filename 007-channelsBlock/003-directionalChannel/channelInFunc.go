package main

import (
	"fmt"
)

func main() {
	// bi - directional channel
	c := make(chan int)
	fmt.Printf("%T\t%v\n", c, "Type of bidirectional channel")
	// channel to set values to the channel
	in := make(chan<- int)
	fmt.Printf("%T\t%v\n", in, "Type of set only values to a channel")
	// channel to recieve values from the channel
	out := make(<-chan int)
	fmt.Printf("%T\t%v\n", out, "Type of recieve only values to a channel")

}
