package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Metadata struct {
	GlobalID       string   `json:"globalID"`
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Titles         []string `json:"titles"`
	TitlesFull     []Title  `json:"titles_full"`
	Link           string   `json:"link"`
	Description    string   `json:"description"`
	Genres         []string `json:"genres"`
	MediaType      string   `json:"mediaType"`
	Rating         float64  `json:"rating"`
	Languages      []string `json:"languages"`
	ContentRating  string   `json:"contentRating"`
	Actors         []string `json:"actors"`
	Trailer        Trailer  `json:"trailer"`
	Characters     []string `json:"characters"`
	Image          string   `json:"image"`
	IsAdult        bool     `json:"isAdult"`
	Type           string   `json:"type"`
	ReleasedDate   string   `json:"releasedDate"`
	EpisodesNumber int      `json:"episodesNumber"`
	Runtime        string   `json:"runtime"`
	ReleaseYears   []int    `json:"-"`
	Keywords       []string `json:"keywords"`
	Backdrop       string   `json:"backdrop"`
}

func (m *Metadata) UnmarshalJSON(d []byte) error {
	type Alias Metadata
	type Aux struct {
		*Alias

		ReleaseYears any `json:"releaseYears"`
	}

	aux := &Aux{
		Alias: (*Alias)(m),
	}

	err := json.Unmarshal(d, aux)
	if err != nil {
		return err
	}

	if aux.ReleaseYears != nil {
		releaseYearsStr, ok := aux.ReleaseYears.(string)
		if !ok {
			return nil
		}

		isRange := strings.Contains(releaseYearsStr, "-")

		var yearStrs []string = []string{releaseYearsStr}
		if isRange {
			yearStrs = strings.Split(aux.ReleaseYears.(string), "-")
		}

		for _, year := range yearStrs {
			yearStr := strings.TrimSpace(year)
			yearInt, err := strconv.Atoi(yearStr)
			if err != nil {
				return err
			}

			m.ReleaseYears = append(m.ReleaseYears, yearInt)
		}

	}

	return nil
}

type Title struct {
	Language string `json:"language"`
	Title    string `json:"title"`
}

type Trailer struct {
	YoutubeID string `json:"youtube_id"`
	FullURL   string `json:"full_url"`
	Thumbnail string `json:"thumbnail"`
}
