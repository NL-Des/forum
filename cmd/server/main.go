package main

import (
	"forum/internal/config"
	"forum/internal/handlers"
	"forum/internal/repositories"
	"forum/internal/services"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()

	//Lancement de la BdD:
	db := config.InitDB()
	defer db.Close()

	//Injection des dépendances:
	//cheminement BdD → repositories → services → handlers
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	handlers.InitHandlers(userService)

	//Routage HTTP:
	//handlers → front
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	fs := http.FileServer(http.Dir("../../internal/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Lancement serveur:
	addr := ":8086"
	log.Printf("Server start → http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal("❌ error running server:", err)
	}
}
