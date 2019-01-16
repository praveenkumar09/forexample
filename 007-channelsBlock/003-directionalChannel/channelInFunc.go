package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("The number of routines first time is ->", runtime.NumGoroutine())
		c <- 42
		c <- 43
		fmt.Println("The value 43 is also added")
	}()
	fmt.Println("The number of routines second time is ->", runtime.NumGoroutine())
	fmt.Println("The value inside the channel is ->", <-c)
	fmt.Println("The value inside the channel is ->", <-c)
	fmt.Println("The number of routines third time is ->", runtime.NumGoroutine())
}
