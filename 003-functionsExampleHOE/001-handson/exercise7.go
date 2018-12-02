package main

import "fmt"

func main() {
	x := func(){
		for i:=0;i<3;i++{
			fmt.Println(i)
		}
	}

	x()
	fmt.Printf("%T\n",x)

	y := x
	y()
	fmt.Printf("%T\n",y)

}

