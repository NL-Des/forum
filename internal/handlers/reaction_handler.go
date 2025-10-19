package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

/*MARK: React
 */
func ReactHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	targetType := r.URL.Query().Get("target_type")
	targetID, _ := strconv.ParseInt(r.URL.Query().Get("target_id"), 10, 64)
	value, _ := strconv.Atoi(r.URL.Query().Get("value"))

	if err := reactionService.React(userID, targetType, targetID, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "reaction saved: user=%d target=%s/%d value=%d\n", userID, targetType, targetID, value)
}

/*MARK: RemoveReaction
 */
func RemoveReactionHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	targetType := r.URL.Query().Get("target_type")
	targetID, _ := strconv.ParseInt(r.URL.Query().Get("target_id"), 10, 64)

	if err := reactionService.RemoveReaction(userID, targetType, targetID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "reaction deleted: user=%d target=%s/%d\n", userID, targetType, targetID)
}

/*MARK: GetCounts
 */
func GetCountsHandler(w http.ResponseWriter, r *http.Request) {
	targetType := r.URL.Query().Get("target_type")
	targetID, _ := strconv.ParseInt(r.URL.Query().Get("target_id"), 10, 64)

	likes, dislikes, err := reactionService.GetReactionCounts(targetType, targetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Target %s/%d: %d likes, %d dislikes\n", targetType, targetID, likes, dislikes)
}
