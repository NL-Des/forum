package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/internal/database"
)

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	//log.Println("Reçu :", title, content)

	if title == "" || content == "" {
		http.Error(w, "Champs requis manquants", http.StatusBadRequest)
		return
	}

	err := database.InsertTopic(title, content, 0)
	if err != nil {
		http.Error(w, "Erreur lors de l'enregistrement :"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	topic, err := database.GetTopicByID(id)
	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusNotFound)
		return
	}
	posts, err := database.GetPostsByTopicID(id)
	if err != nil {
		log.Println("⚠️ Erreur récupération messages :", err)
		posts = []database.Post{}
	}

	data := database.Thread{Topic: topic, Posts: posts}

	tmpl := template.Must(template.ParseFiles("internal/templates/topic.html"))
	tmpl.Execute(w, data)
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	topicID, _ := strconv.Atoi(r.FormValue("topic_id"))
	content := r.FormValue("content")

	err := database.InsertPost(topicID, content, 0)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout du message", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/topic?id=%d", topicID), http.StatusSeeOther)
}
