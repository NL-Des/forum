package repositories

import (
	"forum/internal/config"
	"log"
	"net/http"
)

func RegisterUserInSQL(w http.ResponseWriter, username string, hashedPassword []byte, email string) {

	db := config.InitDB() // Ouverture de la base de données.
	defer db.Close()      // Fermeture à la fin de la fonction.

	// Injection dans la base SQL.
	_ /* position */, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, hashedPassword, email)
	if err != nil {
		http.Error(w, "Error during Register", http.StatusInternalServerError)
		log.Println(err)
	}
	/*
		 	Test pour vérifier les données.
			id, _ := position.LastInsertId()

			fmt.Printf("Username: %s\nPassword: %s\nEmail: %s\nID DataBase :%d\n", username, hashedPassword, email, id)
	*/
}
