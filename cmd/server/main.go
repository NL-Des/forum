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

	//1- Charger le fichier .env
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("❌ error loading .env file")
	}
	//2- Lancement de la BdD:
	db := config.InitDB()
	defer db.Close()

	//3- Injection des dépendances:
	//   cheminement BdD → repositories → services → handlers
	userRepository := repositories.NewUserRepository(db)
	topicPostRepository := repositories.NewTopicPostRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	reactionRepository := repositories.NewReactionRepository(db)
	filterRepository := repositories.NewFilterRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	userService := services.NewUserService(userRepository)
	topicPostService := services.NewTopicPostService(topicPostRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	reactionService := services.NewReactionService(reactionRepository)
	filterService := services.NewFilterService(filterRepository)
	authService := services.NewAuthService(authRepository)

	//4- Récup des Routes HTTP:
	//   handlers → front
	router := handlers.Router(userService, topicPostService, categoryService, reactionService, filterService, authService)

	//5- Lancement serveur:
	addr := os.Getenv("SERVER_PORT")
	GithubClientID := os.Getenv("GITHUB_CLIENT_ID")
	GithubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if addr == "" {
		addr = ":8086" // valeur par défaut en dev, sinon c'est une variable définie dans .env
	}
	if len(GithubClientID) == 0 || len(GithubClientSecret) == 0 {
		log.Fatal("Set GITHUB_CLIENT_* env vars")
	}
	log.Printf("Server start → http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("❌ error trying to run the server: ", err)
	}
}
