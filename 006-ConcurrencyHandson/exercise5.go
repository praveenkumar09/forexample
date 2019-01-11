package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	const endingVal = 100
	var v int64
	var waitGroup sync.WaitGroup
	waitGroup.Add(endingVal)

	for i := 0; i < endingVal; i++ {
		go func() {
			atomic.AddInt64(&v, 1)
			fmt.Println("The value of the v in atomic is :", atomic.LoadInt64(&v))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println("the value showing outside the loop for v is :", v)
}
