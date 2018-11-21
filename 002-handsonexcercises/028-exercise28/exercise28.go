package main

import "fmt"

func main(){
x := []string{"James", "Bond", "Shaken, not stirred"};
y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
xy := [][]string{x,y}
fmt.Println(xy)

for i,v := range xy{
	fmt.Println(i,v)
	for a,b := range v{
		fmt.Println(a,b)
	}
}
}
