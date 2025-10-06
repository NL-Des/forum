package server

import (
	"fmt"
	"forum/server/pages"
	"net/http"
)

func LaunchForum() {
	// Routeur pour les différentes pages du serveur.
	mux := http.NewServeMux()

	// Liste des pages disponibles du forum.
	mux.HandleFunc("/forum", pages.ForumHandler)
	mux.HandleFunc("/register", pages.RegisterHandler)
	mux.HandleFunc("/login", pages.LoginHandler)
	mux.HandleFunc("/logout", pages.LogoutHandler)
	mux.HandleFunc("/subject", pages.SubjectHandler)
	mux.HandleFunc("/newsubject", pages.NewSubjectHandler)

	// Lancement du serveur.
	fmt.Println("Serveur lancé sur http://localhost:8080/forum")
	http.ListenAndServe(":8080", mux)
}
