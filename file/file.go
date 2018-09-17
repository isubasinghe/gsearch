package file

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ScoreData stores the scores and indexes for a single query term
type ScoreData struct {
	Index int
	Score float32
}

// WordData An abstraction layer for word score data
type WordData struct {
	Data map[string][]ScoreData
}

// NewWordData creates the memory for use by the struct
func (w *WordData) NewWordData() {
	w.Data = make(map[string][]ScoreData)
}

// InsertScores inserts the scores for a given word
func (w *WordData) InsertScores(word string, scoreData []ScoreData) {
	w.Data[word] = scoreData
}

// GetScores gets the scores for a given word
func (w *WordData) GetScores(word string) []ScoreData {
	return w.Data[word]
}

// SaveToGOB save the data structure
func (w *WordData) SaveToGOB(fname string) {
	writeGob(fname, *w)
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// GetIndexedURLs returns the indexed urls
func GetIndexedURLs(fpath string) ([]string, error) {
	var urlList []string
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	var index int
	var url string
	for {
		out, err := fmt.Fscanf(file, "%d %s\n", &index, &url)
		if err != nil {
			break
		}
		if out != 2 {
			return nil, errors.New("Invalid data format")
		}
		urlList = append(urlList, url)

	}
	return urlList, nil
}

// GetIndexedURLsGOB reads the gobified document data
func GetIndexedURLsGOB(fpath string) ([]string, error) {
	var urlList = new([]string)
	err := readGob("./documents.gob", urlList)
	if err != nil {
		return nil, err
	}
	return *urlList, nil
}

func readFile(fname string) ([]ScoreData, error) {
	var scoreData []ScoreData
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	for {
		var score float32
		var index int
		out, err := fmt.Fscanf(file, "%d %f\n", &index, &score)
		if err != nil {
			break
		}
		if out != 2 {
			return nil, errors.New("Unable to understand file format")
		}
		fileScore := ScoreData{Index: index, Score: score}
		scoreData = append(scoreData, fileScore)
	}
	sort.Slice(scoreData, func(i, j int) bool {
		return scoreData[i].Score > scoreData[j].Score
	})
	return scoreData, nil
}

// LoadWordDataGOB loads the data gob from a file
func LoadWordDataGOB(path string) (WordData, error) {
	var wordData = new(WordData)
	wordData.NewWordData()
	err := readGob(path, wordData)
	if err != nil {
		return *wordData, err
	}
	return *wordData, nil
}

// LoadData loads the word data
func LoadData(path string) (WordData, error) {
	var wordData WordData
	_ = wordData
	wordData.NewWordData()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return wordData, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if strings.Contains(file.Name(), ".txt") {
				fname := strings.Replace(file.Name(), ".txt", "", -1)
				scores, err := readFile("./data/" + file.Name())
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println("Inserted:", file.Name())
				wordData.InsertScores(fname, scores)
			}
		}
	}
	return wordData, nil
}
