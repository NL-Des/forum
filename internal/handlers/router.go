package handlers

import (
	"forum/internal/domain"
	"net/http"
)

func Router(userService domain.UserService, topicPostService domain.TopicPostService, categoryService domain.CategoryService, reactionService domain.ReactionService, filterService domain.FilterService, authService domain.AuthService) http.Handler {
	InitHandlers(userService, topicPostService, categoryService, reactionService, filterService, authService)

	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/thread", ThreadHandler)
	mux.HandleFunc("/login", AuthenticateHandler)
	mux.HandleFunc("/logout", LogoutHandler)
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/create-topic", CreateTopicHandler)
	mux.HandleFunc("/topic", TopicHandler)
	mux.HandleFunc("/add-post", AddPostHandler)
	mux.HandleFunc("/react", ReactHandler)
	mux.HandleFunc("/remove-reaction", RemoveReactionHandler)
	mux.HandleFunc("/filter", FilterTopicByUser)
	mux.HandleFunc("/api/login-gh/", githubLoginHandler)
	mux.HandleFunc("/api/github/callback/", githubCallbackHandler)

	fs := http.FileServer(http.Dir("internal/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}
