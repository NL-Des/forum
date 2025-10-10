package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("../../internal/templates/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Erreur :", http.StatusNotFound)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Println("Erreur template :", err)
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
