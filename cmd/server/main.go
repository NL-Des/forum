package main

import (
	"forum/internal/config"
	"forum/internal/handlers"
	"forum/internal/repositories"
	"forum/internal/services"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/*pw := "motdepasse"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	fmt.Println(string(hash))*/

	// Charger le fichier .env
	/* 	errEnv := godotenv.Load()
	   	if errEnv != nil {
	   		log.Fatal("❌ error loading .env file")
	   	} */

	//Lancement de la BdD:
	//Injection des dépendances:
	//cheminement BdD → repositories → services → handlers
	/*
		userRepo := repositories.NewUserRepository(db)
		userService := services.NewUserService(userRepo)
		handlers.InitHandlers(userService)*/

	db := config.InitDB()
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	handlers.InitHandlers(userService)

	mux := http.NewServeMux()

	//Routage HTTP:
	//handlers → front
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/login", handlers.Authenticate)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/create-topic", handlers.CreateTopicHandler)
	mux.HandleFunc("/topic", handlers.TopicHandler)
	mux.HandleFunc("/add-post", handlers.AddPostHandler)

	fs := http.FileServer(http.Dir("../../internal/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Lancement serveur:
	addr := os.Getenv("SERVER_PORT")
	if addr == "" {
		addr = ":8086" // valeur par défaut en dev, sinon c'est une variable définie dans .env
	}
	log.Printf("Server start → http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal("❌ error trying to run the server: ", err)
	}
}
