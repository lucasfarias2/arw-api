package handlers

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Error 404: Not found"))
	if err != nil {
		return
	}
}
