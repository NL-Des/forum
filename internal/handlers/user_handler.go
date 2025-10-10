package handlers

import (
	"forum/internal/services"
	"html/template"
	"net/http"
	"path/filepath"
)

var userService services.UserService

func InitHandlers(us services.UserService) {
	userService = us
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("internal", "templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur interne : "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "register.html", nil)
		return
	}

	// Méthode POST : récupération du formulaire
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := userService.Register(username, email, password)
	if err != nil {
		// Réaffiche le formulaire avec une erreur
		renderTemplate(w, "register.html", map[string]string{
			"Error": err.Error(),
		})
		return
	}

	// Enregistrement réussi, redirection vers la page d’accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
