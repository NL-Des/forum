package handlers

import (
	"forum/internal/domain"
	"log"
	"net/http"
)

type Datas struct {
	Topics     []domain.Topic
	IsLoggedIn bool
}

/*var tpl = template.Must(template.ParseFiles("internal/templates/home.html"))*/

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "❌ not found", http.StatusNotFound)
		return
	}

	topics, err := topicPostService.GetAllTopics()
	if err != nil {
		log.Println("❌ error fetching topics:", err)
		http.Error(w, "❌ error fetching topics", http.StatusInternalServerError)
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

	datas := Datas{
		Topics:     topics,
		IsLoggedIn: isLoggedIn,
	}
	log.Printf("Nombre de topics: %d\n", len(topics))
	RenderTemplate(w, "home.html", datas)
}
