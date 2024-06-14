package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jjalbuenacabuyao/dictionary/fetch"
	"github.com/jjalbuenacabuyao/dictionary/fmtresponse"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../index.html"))

	tmpl.Execute(w, nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	res, err := fetch.FetchData(r.PostFormValue("search"))
	if err != nil {
		log.Fatal(err.Error())
	}

	minimizedData := fmtresponse.FormatApiResponse(res)

	tmpl := template.Must(template.ParseFiles("templates/template.html"))

	tmpl.Execute(w, &minimizedData)
}