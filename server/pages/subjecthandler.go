package pages

import "net/http"

func SubjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Subject."))
}
