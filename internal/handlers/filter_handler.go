package handlers

import (
	"fmt"
	"forum/internal/domain"
	"log"
	"net/http"
)

type Datas struct {
	Topics     []domain.Topic
	Categories []domain.Category
	IsLoggedIn bool
}

func FilterTopicByUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "❌ unauthorized method", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur de parsing du formulaire", http.StatusBadRequest)
		fmt.Println("err de parse")
		return
	}

	cookie, err := r.Cookie("session_token")
	var isLoggedIn bool
	var topics []domain.Topic
	var categories []domain.Category
	if err == nil {
		// Vérifie si le token correspond à un utilisateur connecté
		user, _ := userService.Home(cookie.Value)
		if user != nil {
			isLoggedIn = true
		}
	}

	filter := r.FormValue("filters")
	FormCategory := r.FormValue("categories")
	user, err := userService.Home(cookie.Value)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}
	fmt.Println(filter)
	if filter == "messages" && FormCategory == "" {
		topics, err = topicPostService.FilterTopic(int(user.ID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if topics == nil {
			fmt.Println("utilisateur sans topics")
		}
		for i := range topics {
			categories, err = categoryService.GetCategoriesByTopicID(topics[i].ID)
			if err != nil {
				log.Println("❌ error fetching categories:", err)
				http.Error(w, "❌ error fetching categories", http.StatusInternalServerError)
				return
			}
		}
	} else if filter == "messages" && FormCategory != "" {

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	/*topics, err := topicPostService.GetAllTopics()
	if err != nil {
		log.Println("❌ error fetching topics:", err)
		http.Error(w, "❌ error fetching topics", http.StatusInternalServerError)
		return
	}*/
	for i := range topics {
		topics[i].Categories, err = categoryService.GetCategoriesByTopicID(topics[i].ID)
		if err != nil {
			log.Println("❌ error fetching categories:", err)
			http.Error(w, "❌ error fetching categories", http.StatusInternalServerError)
			return
		}
	}
	/*categories, err = categoryService.GetAllCategories()
	if err != nil {
		log.Println("❌ error fetching categories:", err)
		http.Error(w, "❌ error fetching categories", http.StatusInternalServerError)
		return
	}*/

	datas := Datas{
		Topics:     topics,
		IsLoggedIn: isLoggedIn,
		Categories: categories,
	}
	log.Printf("Nombre de topics: %d\n", len(topics))
	RenderTemplate(w, "home.html", datas)
}
