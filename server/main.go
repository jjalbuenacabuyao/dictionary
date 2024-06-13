package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type ApiResponse []struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string   `json:"definition"`
			Example    string   `json:"example"`
		} `json:"definitions"`
		Synonyms   []string `json:"synonyms"`
		Antonyms   []string `json:"antonyms"`
	} `json:"meanings"`
	Source []string `json:"sourceUrls"`
}

type MinimizedApiResponse struct {
	Word          string
	PhoneticText  string
	PhoneticAudio string
	Meanings      []struct {
		PartOfSpeech string
		Definitions  []struct {
			Definition string
			Example    string
		}
		Synonyms   []string
		Antonyms   []string
	}
	Source string
}

func FetchData(word string) (ApiResponse, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	response, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ApiResponse{}, err
	}

	var data ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return ApiResponse{}, err
	}

	return data, nil
}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../index.html"))

	tmpl.Execute(w, nil)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	res, err := FetchData(r.PostFormValue("search"))
	if err != nil {
		log.Fatal(err.Error())
	}

	var minimizedData MinimizedApiResponse
	for _, v := range res {
		minimizedData.Word = v.Word
		minimizedData.PhoneticText = v.Phonetics[0].Text

		for _, p := range v.Phonetics {
			if p.Audio != "" {
				minimizedData.PhoneticAudio = p.Audio
				break
			}
		}

		minimizedData.Meanings = []struct {
			PartOfSpeech string
			Definitions  []struct {
				Definition string
				Example    string
			}
			Synonyms   []string
			Antonyms   []string
		}(v.Meanings)
		minimizedData.Source = v.Source[0]
	}

	tmpl := template.Must(template.ParseFiles("templates/template.html"))

	tmpl.Execute(w, minimizedData)
}

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleHomepage)
	http.HandleFunc("/search", handleSearch)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
