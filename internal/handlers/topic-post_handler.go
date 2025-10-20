package handlers

import (
	"fmt"
	"html/template"
	"log"
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

	// parser le formulaire
	if err := r.ParseForm(); err != nil {
		http.Error(w, "❌ cannot parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	/*log.Println("Form values:", r.Form)*/

	content := r.FormValue("content")
	topicID, _ := strconv.Atoi(r.FormValue("topic_id"))
	userID := 3
	//userID := r.Context().Value("userID").(int)

	err2 := topicPostService.AddPost(topicID, content, userID)
	if err2 != nil {
		log.Println("❌ AddPost error:", err2)
		http.Error(w, "❌ error inserting post:"+err2.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("✅ Post inserted, redirecting to thread")
	http.Redirect(w, r, fmt.Sprintf("/thread?id=%d", topicID), http.StatusSeeOther)
}
