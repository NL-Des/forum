package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("./internal/templates/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Erreur :", http.StatusNotFound)
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

	data := map[string]interface{}{
		"IsLoggedIn": isLoggedIn,
	}

	if err := tpl.Execute(w, data); err != nil {
		log.Println("Erreur template :", err)
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
