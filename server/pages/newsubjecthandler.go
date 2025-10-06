package pages

import "net/http"

func NewSubjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Newsubject."))
}
