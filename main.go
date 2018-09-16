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
	f.LoadData("./data")
}
