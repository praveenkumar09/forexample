package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		dob   string
		age   int
	}{
		first: "Praveen",
		last:  "Kumar",
		dob:   "12/12/1991",
		age:   26,
	}
	fmt.Println(p1)
}
