package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go foo1(c)

	bar1(c)
	//fmt.Println("Code about to exit")

}

func foo1(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("The value about to add in the channel is =>", i)
		c <- i
	}
	close(c)
}

func bar1(c <-chan int){
	for v := range c {
		fmt.Println("The value about to recieve is =>", v)
	}
}
