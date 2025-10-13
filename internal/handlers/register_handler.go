package handlers

import (
	"forum/internal/repositories"
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

	// Fonction pour chiffrer le mot de passe.
	hashedPassword := repositories.HashPassword(password)
	// Fonction pour enregistrer les données du nouvel utilisateur.
	repositories.RegisterUserInSQL(w, username, hashedPassword, email)

	// Enregistrement réussi, redirection vers la page d’accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
