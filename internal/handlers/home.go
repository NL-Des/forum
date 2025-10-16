package handlers

import (
	"forum/internal/database"
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("internal/templates/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Erreur :", http.StatusNotFound)
		return
	}

	topics, err := database.GetAllTopics()
	if err != nil {
		log.Println("Erreur récupération sujets :", err)
		http.Error(w, "Erreur lors du chargement des sujets", http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("session_token")
	var isLoggedIn bool

	if err == nil {
		// Vérifie si le token correspond à un utilisateur connecté
		user, _ := userService.Home(cookie.Value)
		if &user != nil {
			isLoggedIn = true
		}
	}

	datas := database.Datas{
		Topics:     topics,
		IsLoggedIn: isLoggedIn,
	}
	renderTemplate(w, "home.html", datas)
	/*if err := tpl.Execute(w, datas); err != nil {
		log.Println("Erreur template :", err)
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}*/

	/* if err := tpl.Execute(w, nil); err != nil {
		log.Println("Erreur template :", err)
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	} */
}
