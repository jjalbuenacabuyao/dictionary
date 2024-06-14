package main

import (
	"log"
	"net/http"

	"github.com/jjalbuenacabuyao/dictionary/handler"
)

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.Homepage)
	http.HandleFunc("/search", handler.Search)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
