package controller

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/test2/models"
)

func sorting(s string) string {
	text := strings.ToLower(s)
	newText := strings.Split(text, "")
	sort.Strings(newText)

	return strings.Join(newText, "")
}

// HandlerSorting handle coming request
func HandlerSorting(w http.ResponseWriter, r *http.Request) {

	var word models.Word
	json.NewDecoder(r.Body).Decode(&word)
	resolve := sorting(word.Text)

	word.Text = resolve
	json.NewEncoder(w).Encode(word)
}
