package handler

import (
	"net/http"
)

func ApplicationHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Application is healthy"))
}

