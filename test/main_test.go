package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/test2/models"
)

func Test(t *testing.T) {

	word := models.Word{Text: "Awan"}
	body := new(bytes.Buffer)

	json.NewEncoder(body).Encode(word)
	req, err := http.NewRequest("POST", "http://localhost", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	client := &http.Client{}
	res, err := client.Do(req)

	var w models.Word
	json.NewDecoder(res.Body).Decode(&w)

	if w.Text != "aanw" {
		t.Errorf("Result must be sorted character a-z")
	}
}
