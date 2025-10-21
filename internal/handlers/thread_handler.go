package handlers

import (
	"forum/internal/domain"
	"html/template"
	"net/http"
	"strconv"
)

type ThreadData struct {
	Topic      domain.Topic
	Posts      []domain.Post
	Categories []domain.Category
}

func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "❌ missing ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ invalid ID", http.StatusBadRequest)
		return
	}

	// Récupération du thread (topic + posts)
	thread, err := topicPostService.GetThreadByID(id)
	if err != nil {
		http.Error(w, "❌ topic not found", http.StatusNotFound)
		return
	}

	thread.Categories, err = categoryService.GetCategoriesByTopicID(id)
	if err != nil {
		http.Error(w, "❌ categories not found", http.StatusNotFound)
		return
	}

	// Rendu du template thread.html
	tmpl := template.Must(template.ParseFiles("internal/templates/thread.html"))
	err = tmpl.Execute(w, thread)
	if err != nil {
		http.Error(w, "❌ template error: "+err.Error(), http.StatusInternalServerError)
	}
}
