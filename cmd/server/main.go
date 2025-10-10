package main

import (
	"forum/internal/config"
	"forum/internal/handlers"
	"forum/internal/repositories"
	"forum/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/*pw := "motdepasse"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	fmt.Println(string(hash))*/

	mux := http.NewServeMux()

	// Charger le fichier .env
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("❌ error loading .env file")
	}

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
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	fs := http.FileServer(http.Dir("templates/assets"))
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
