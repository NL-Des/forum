package handlers

import (
	"forum/internal/domain"
	"net/http"
)

func Router(userService domain.UserService, topicPostService domain.TopicPostService, categoryService domain.CategoryService, reactionService domain.ReactionService, filterService domain.FilterService) http.Handler {
	InitHandlers(userService, topicPostService, categoryService, reactionService, filterService)

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

	fs := http.FileServer(http.Dir("internal/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}
