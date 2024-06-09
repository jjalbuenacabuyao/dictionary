package main

import (
	"encoding/json"

	"html/template"
	"io"
	"log"
	"net/http"
)

type Phonetic struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type Phonetics struct {
	Phonetic []Phonetic
}

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    string   `json:"example"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
}

type ApiResponse struct {
	Word      string      `json:"word"`
	Phonetics []Phonetics `json:"phonetics"`
	Meanings  []Meaning   `json:"meanings"`
	Source    []string    `json:"sourceUrls"`
}

func FetchData(word string) ([]ApiResponse, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	response, err := http.Get(url)
	if err != nil {
		return []ApiResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []ApiResponse{}, err
	}

	var data []ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return []ApiResponse{}, err
	}

	return data, nil
}

func Respond(w http.ResponseWriter, r *http.Request) {
	res, err := FetchData(r.PostFormValue("search"))
	if err != nil {
		log.Fatal(err.Error())
	}

	tmpl := template.Must(template.New("temp").Parse(`
		<p>{{.Word}}</p>
		
		<pre>{{.Meanings}}</pre>
	`))

	tmpl.Execute(w, res[0])
}

func main() {
	http.HandleFunc("/search", Respond)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
