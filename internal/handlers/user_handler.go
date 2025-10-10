package handlers

import (
	"forum/internal/database"
	"forum/internal/services"
	"html/template"
	"log"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("internal/templates/home.html"))
	if r.URL.Path != "/" {
		http.Error(w, "Erreur :", http.StatusNotFound)
		return
	}

	topics, err := database.GetAllTopics()
	if err != nil {
		log.Println("Erreur récupération sujets :", err)
		http.Error(w, "Erreur lors du chargement des sujets", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, topics); err != nil {
		log.Println("Erreur template :", err)
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
