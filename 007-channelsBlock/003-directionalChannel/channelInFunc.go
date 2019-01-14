package main

import (
	"fmt"
	"runtime"
)

//var wg sync.WaitGroup

func main(){
	c := make(chan int)
	fmt.Println(c)
	//wg.Add(2)
	// send to the foo function
	go foo(c)

	//recieve from the bar function
	bar(c)

	// wg.Wait()
	fmt.Println("The code about to exit")
}

func foo(c chan<- int){
	fmt.Println("GORout's",runtime.NumGoroutine())
	c <- 42
	//wg.Done()
}

func bar(c <-chan int){
	fmt.Println("GORout's",runtime.NumGoroutine())
	fmt.Println("the value we got from the channel is :",<-c)
	//wg.Done()
}
