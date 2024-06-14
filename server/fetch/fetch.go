package fetch

import (
	"encoding/json"
	"io"
	"net/http"

	rst "github.com/jjalbuenacabuyao/dictionary/responsetype"
)

func FetchData(word string) (*rst.ApiResponse, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	response, err := http.Get(url)
	if err != nil {
		return &rst.ApiResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &rst.ApiResponse{}, err
	}

	var data rst.ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return &rst.ApiResponse{}, err
	}

	return &data, nil
}