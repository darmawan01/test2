package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
)

// Word models
type Word struct {
	Text string `json:"text"`
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func sorting(s string) string {
	text := strings.ToLower(s)
	newText := strings.Split(text, "")
	sort.Strings(newText)

	return strings.Join(newText, "")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var word Word
	json.NewDecoder(r.Body).Decode(&word)
	resolve := sorting(word.Text)

	word.Text = resolve
	json.NewEncoder(w).Encode(word)
}

func main() {

	http.HandleFunc("/", logging(welcome))

	http.ListenAndServe(":80", nil)
}
