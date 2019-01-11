package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const endingVal = 100
	initialVal := 0
	var v int
	var waitGroup sync.WaitGroup
	var mu sync.Mutex
	waitGroup.Add(endingVal)

	fmt.Println("The number of CPUs is :", runtime.NumCPU())
	fmt.Println("The number of GoRoutines starting is :", runtime.NumGoroutine())

	for i := 0; i < endingVal; i++ {
		go func() {
			mu.Lock()
			v = initialVal
			runtime.Gosched()
			initialVal++
			v = initialVal
			fmt.Println("the value showing inside the loop for v is :", v)
			fmt.Println("The number of GoRoutines inside is :", runtime.NumGoroutine())
			waitGroup.Done()
			mu.Unlock()
		}()
	}

	waitGroup.Wait()
	fmt.Println("The number of GoRoutines outside is :", runtime.NumGoroutine())
	fmt.Println("the value showing outside the loop for v is :", v)
}
