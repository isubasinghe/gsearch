package main

import (
	"fmt"
	"net/http"
	"strings"

	f "./file"
	s "./search"
)

var (
	search s.Search
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	queries, ok := r.URL.Query()["query"]
	if !ok || len(queries) < 1 {
		return
	}
	query := strings.Split(queries[0], " ")
	url := search.Search(query)
	w.Write([]byte(url))
}

func main() {
	urls, err := f.GetIndexedURLsGOB("./documents.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	wordData, err := f.LoadWordDataGOB("./data.gob")
	if err != nil {
		fmt.Println(err)
		return
	}

	search.NewSearch(wordData, urls)

	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":8080", nil)

}
