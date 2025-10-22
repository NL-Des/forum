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
	IsLoggedIn bool
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

    cookie, err := r.Cookie("session_token")
    var isLoggedIn bool

    if err == nil {
        // Vérifie si le token correspond à un utilisateur connecté
        user, _ := userService.Home(cookie.Value)
        if user != nil {
            isLoggedIn = true
        }
    }

    // Préparation des données pour le template
    datas := ThreadData{
        Topic:      thread.Topic,
        Posts:      thread.Posts,
        IsLoggedIn: isLoggedIn,
    }

    // Rendu du template avec les bonnes données
    tmpl := template.Must(template.ParseFiles("internal/templates/thread.html"))
    err = tmpl.Execute(w, datas)
    if err != nil {
        http.Error(w, "❌ template error: "+err.Error(), http.StatusInternalServerError)
    }
}
