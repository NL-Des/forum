package handlers

import (
	"forum/internal/domain"
	"net/http"
)

func Router(userService domain.UserService, topicPostService domain.TopicPostService) http.Handler {
	InitHandlers(userService, topicPostService)

	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/login", Authenticate)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/create-topic", CreateTopicHandler)
	mux.HandleFunc("/topic", TopicHandler)
	mux.HandleFunc("/add-post", AddPostHandler)

	fs := http.FileServer(http.Dir("internal/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}
