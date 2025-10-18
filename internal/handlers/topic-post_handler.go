package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

/*MARK: CreateTopic
 */
func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "❌ unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	//log.Println("Reçu :", title, content)

	if title == "" || content == "" {
		http.Error(w, "❌ missing required fields", http.StatusBadRequest)
		return
	}

	err := topicPostService.CreateTopic(title, content, 0)
	if err != nil {
		http.Error(w, "❌ error inserting topic"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// affichage d'un sujet et ses messages
/*MARK: Topic+Posts
 */
func TopicHandler(w http.ResponseWriter, r *http.Request) {
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

	thread, err := topicPostService.GetThreadByID(id)
	if err != nil {
		http.Error(w, "❌ topic not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("internal/templates/topic.html"))
	tmpl.Execute(w, thread)
}

/*MARK: AddPost
 */
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "❌ unauthorized method", http.StatusMethodNotAllowed)
		return
	}

	topicID, _ := strconv.Atoi(r.FormValue("topic_id"))
	content := r.FormValue("content")

	err := topicPostService.AddPost(topicID, content, 0)
	if err != nil {
		http.Error(w, "❌ error inserting post:"+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/topic?id=%d", topicID), http.StatusSeeOther)
}
