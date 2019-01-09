package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "praveenkumar"
	pb, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pb))

	err = bcrypt.CompareHashAndPassword(pb, []byte(password))
	if err == nil {
		fmt.Println("Both hash and password matches")
	} else {
		fmt.Println("Hash and password doesnt match")
	}
}
