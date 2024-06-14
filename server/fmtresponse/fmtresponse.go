package fmtresponse

import rst "github.com/jjalbuenacabuyao/dictionary/responsetype"

func FormatApiResponse(res *rst.ApiResponse) (*rst.MinimizedApiResponse) {
	var minimizedData rst.MinimizedApiResponse

	for _, v := range *res {
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

	return &minimizedData
}