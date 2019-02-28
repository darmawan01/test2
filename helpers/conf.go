package helpers

import (
	"log"
	"net/http"
)

// Logging log request
func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// EnableCors set cor enable
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}
