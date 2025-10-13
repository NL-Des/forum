package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	test, err := bcrypt.GenerateFromPassword([]byte("hashedpassword13"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(test))
}
