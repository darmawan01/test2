package main

import (
	"net/http"

	ws "github.com/test2/handlers"
	cf "github.com/test2/helpers"
)

func main() {

	http.HandleFunc("/", cf.Logging(ws.HandlerSorting))
	http.ListenAndServe(":80", nil)
}
