package pages

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "forum.html")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
	}
	template.Execute(w, nil)
}
