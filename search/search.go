package search

import (
	"log"

	f "../file"
)

type Search struct {
	wordData f.WordData
	urlList  []string
}

func NewSearchEngine() Search {
	return Search{}
}

func (s *Search) NewSearch(wordData f.WordData, urlList []string) {
	s.wordData = wordData
	s.urlList = urlList
}

func (s *Search) Search(query []string) string {
	scoresHolder := make([]float64, 131563)
	for i, val := range query {
		if i > 255 {
			break
		}
		scores := s.wordData.GetScores(val)
		for _, score := range scores {
			if score.Index > 131563 {
				log.Println("Weird Index at", score.Index, val)
				return ""
			}
			scoresHolder[score.Index] += score.Score
		}
	}

	bestIndex := -1
	bestScore := 0.0
	for i, val := range scoresHolder {
		if val > bestScore {
			bestScore = val
			bestIndex = i
		}
	}
	if bestIndex != -1 {
		return s.urlList[bestIndex]
	}
	return ""
}
