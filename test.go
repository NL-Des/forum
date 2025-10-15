package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	test, err := bcrypt.GenerateFromPassword([]byte("76hashedpassword"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(test))
}
