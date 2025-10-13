package repositories

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return hashedPassword
}
