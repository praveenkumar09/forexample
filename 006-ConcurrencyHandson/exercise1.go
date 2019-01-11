package main

import (
	"fmt"
	"runtime"
	"sync"
)

const firstOneToExec = 100
const secondOneToExec = 200

var waitGroup sync.WaitGroup

func main() {
	fmt.Println("CPu's ", runtime.NumCPU())
	waitGroup.Add(2)
	fmt.Println("Go Routines ", runtime.NumGoroutine())
	go firstOneToExecute()
	go secondOneToExecute()
	waitGroup.Wait()
}

func firstOneToExecute() {
	for i := 0; i <= firstOneToExec; i++ {
		fmt.Println("From first One To Execute ==>", i)
	}
	fmt.Println("Go Routines ", runtime.NumGoroutine())
	waitGroup.Done()
}

func secondOneToExecute() {
	for i := 101; i <= secondOneToExec; i++ {
		fmt.Println("From second One To Execute ==>", i)
	}
	fmt.Println("Go Routines ", runtime.NumGoroutine())
	waitGroup.Done()
}
