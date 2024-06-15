package fmtresponse

import (
	rst "github.com/jjalbuenacabuyao/dictionary/responsetype"
)

func FormatApiResponse(r *rst.ApiResponse) *rst.MinimizedApiResponse {
	resp := *r
	var minimizedData rst.MinimizedApiResponse

	minimizedData.Word = resp[0].Word
	minimizedData.Source = resp[0].Source[0]

	for i, v := range resp {
		if minimizedData.PhoneticText == "" {
			minimizedData.PhoneticText = v.Phonetics[i].Text
		}

		if minimizedData.PhoneticAudio == "" {
			for _, p := range v.Phonetics {
				if p.Audio != "" {
					minimizedData.PhoneticAudio = p.Audio
					break
				}
			}
		}

		for _, m := range v.Meanings {
			minimizedData.Meanings = append(minimizedData.Meanings, rst.Meaning(m))
		}
	}

	return &minimizedData
}
