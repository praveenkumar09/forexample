package main

import "fmt"

func main(){
	x := [10]int{1,2,3,4,5,6,7,8,9,0}
	for i,v:= range x{
		fmt.Println(i,v)
		fmt.Printf("%T\n",v)
	}
}
