package responsetype

type ApiResponse []struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []string `json:"antonyms"`
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
		Synonyms []string
		Antonyms []string
	}
	Source string
}