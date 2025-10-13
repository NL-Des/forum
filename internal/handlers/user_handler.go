package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"forum/internal/services"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var userService services.UserService

func InitHandlers(us services.UserService) {
	userService = us
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("internal/templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur interne : "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testttetstt")
	if r.Method == http.MethodGet {
		// Affiche le formulaire logout via home.html
		renderTemplate(w, "home.html", nil)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err == nil {
		err := userService.Logout(cookie.Value)
		if err != nil {
			fmt.Println(err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now().Add(-1 * time.Hour),
		})
	} else {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Authenticate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Affiche le formulaire login via home.html
		renderTemplate(w, "home.html", nil)
		return
	}

	// Méthode POST : récupération du formulaire
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := userService.Authenticate(email, password)
	if err != nil {
		// Réaffiche home.html avec message d'erreur
		renderTemplate(w, "home.html", map[string]string{
			"Error": err.Error(),
		})
		return
	}

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("erreur : impossible de générer le token ")
	}

	sessionToken := base64.URLEncoding.EncodeToString(bytes)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
	})

	userService.TokenLogIn(sessionToken, email)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "register.html", nil)
		return
	}

	// Méthode POST : récupération du formulaire
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := userService.Register(username, email, password)
	if err != nil {
		// Réaffiche le formulaire avec une erreur
		renderTemplate(w, "register.html", map[string]string{
			"Error": err.Error(),
		})
		return
	}

	// Enregistrement réussi, redirection vers la page d’accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
