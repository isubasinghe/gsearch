package main

import (
	"fmt"

	f "./file"
)

func main() {
	urls, err := f.GetIndexedURLsGOB("./documents.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = urls
	wordData, err := f.LoadDataGob("./data.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	scores := wordData.GetScores("a")
	for i, val := range scores {
		if i > 10 {
			break
		}
		fmt.Println(val.Score, val.Index)
	}
	_ = wordData
}
