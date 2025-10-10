package handlers

import (
	"forum/internal/domain"
	"html/template"
	"net/http"
	"path/filepath"
)

var userService domain.UserService
var templates *template.Template

func InitHandlers(us domain.UserService) {
	userService = us

	// Précharger tous les templates une seule fois
	var err error
	templates, err = template.ParseGlob(filepath.Join("internal", "templates", "*.html"))
	if err != nil {
		panic("❌ error parsing templates: " + err.Error())
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	/* err := templates.ExecuteTemplate(w, tmpl, data) */
	tmplPath := filepath.Join("../../internal/templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "❌ internal error: "+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "❌ error invalid form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := userService.Register(username, email, password)
	if err != nil {
		// Réaffiche le formulaire avec une erreur
		renderTemplate(w, "register.html", map[string]string{
			"❌ error": err.Error(),
		})
		return
	}

	// Enregistrement réussi, redirection vers la page d’accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
