package handlers

import (
	"fmt"
	"forum/internal/config"
	"log"
	"net/http"
)

// Page d'enregistrement de l'utilisateur.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		renderTemplate(w, "register.html", nil)
		return
	}

	err := r.ParseForm() // Si il y a des erreurs (Format,..).
	if err != nil {
		http.Error(w, "Error during the traitement of the formulaire", http.StatusBadRequest)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	db := config.InitDB() // Ouverture de la base de données.
	defer db.Close()      // Fermeture à la fin de la fonction.

	// Injection dans la base SQL.
	position, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, password, email)
	if err != nil {
		http.Error(w, "Error during Register", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	id, _ := position.LastInsertId()

	fmt.Printf("Username: %s\nPassword: %s\nEmail: %s\nID DataBase :%s\n", username, password, email, id)

	// Enregistrement réussi, redirection vers la page d’accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
