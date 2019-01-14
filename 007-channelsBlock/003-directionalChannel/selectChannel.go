package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go setValuesInChannels(even, odd, quit)

	recieveValuesFromchannels(even, odd, quit)

	fmt.Println("The code existed")

}

func recieveValuesFromchannels(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("The value from the even channel is :", v)
		case v := <-o:
			fmt.Println("The value from the odd channel is :", v)
		case v := <-q:
			fmt.Println("The value from the quit channel is :", v)
			return
		}
	}
}

func setValuesInChannels(e, o, q chan<- int) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
}
