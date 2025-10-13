package repositories

import (
	"fmt"
	"forum/internal/config"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return hashedPassword
}

func RegisterUserInSQL(w http.ResponseWriter, username string, hashedPassword []byte, email string) {

	db := config.InitDB() // Ouverture de la base de données.
	defer db.Close()      // Fermeture à la fin de la fonction.

	// Injection dans la base SQL.
	/* position */_ , err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, hashedPassword, email)
	if err != nil {
		http.Error(w, "Error during Register", http.StatusInternalServerError)
		log.Println(err)
	}
/* 	Test pour vérifier les données.
	id, _ := position.LastInsertId()

	fmt.Printf("Username: %s\nPassword: %s\nEmail: %s\nID DataBase :%d\n", username, hashedPassword, email, id) */
}
