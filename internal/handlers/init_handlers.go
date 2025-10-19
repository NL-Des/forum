package handlers

import (
	"fmt"
	"forum/internal/domain"
	"html/template"
	"net/http"
	"path/filepath"
)

var userService domain.UserService
var topicPostService domain.TopicPostService
var templates *template.Template

func InitHandlers(us domain.UserService, tps domain.TopicPostService) {
	userService = us
	topicPostService = tps

	// Précharger tous les templates une seule fois
	var err error
	templates, err = template.ParseGlob(filepath.Join("internal", "templates", "*.html"))
	if err != nil {
		panic("❌ error parsing templates: " + err.Error())
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		fmt.Printf("template error: %v\n", err)
		//http.Error(w, "❌ internal error: "+err.Error(), http.StatusInternalServerError)
		w.Write([]byte("❌ internal error: " + err.Error()))
	}
}
