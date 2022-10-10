package accenter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	wiki "example.com/accenter/internal/wiki"
)

// Load a dataset and attach a weight to each word.
func LoadWeiDataset() []wiki.WeiWikiRecord {

	wikiPath := findDataset("wikiRecords")
	wikiRecords := loadWikiRecords(wikiPath)

	infoPath := findDataset("infoRecords")
	loadInfoWords(infoPath)

	wei_wrecords := make([]wiki.WeiWikiRecord, len(wikiRecords))
	for i, wr := range wikiRecords {
		wei_wrecords[i] = wiki.NewWeiWikiRecord(wr, 1)
	}

	return wei_wrecords
}

func findDataset(which string) string {
	switch which {

	case "wikiRecords":
		dataPath := filepath.Join("..", "..", "dataset", "wiki01.jsonl")
		dataPath, err := filepath.Abs(dataPath)
		if err != nil {
			// if we cannot even built the path of the wiki data file,
			// we have a big problem
			log.Fatal(err)
		}
		fmt.Printf("Loading %s from %s\n", which, dataPath)
		return dataPath

	case "infoRecords":
		dataPath := filepath.Join("..", "..", "dataset", "info01.json")
		dataPath, err := filepath.Abs(dataPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Loading %s from %s\n", which, dataPath)
		return dataPath

	default:
		// MAYBE we should return "", err ?
		return ""
	}
}

// read a file line by line
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
func loadWikiRecords(wikiPath string) []wiki.WikiRecord {
	file, err := os.Open(wikiPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wikiRecords []wiki.WikiRecord

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		// var result map[string]interface{}
		// json.Unmarshal([]byte(line), &result)

		var result_struct wiki.WikiRecord
		json.Unmarshal([]byte(line), &result_struct)

		wikiRecords = append(wikiRecords, result_struct)

		// fmt.Println(line)
		// fmt.Println(result)
		// fmt.Printf("%+v\n\n", result_struct)

		// break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wikiRecords
}

func loadInfoWords(infoPath string) map[string]wiki.InfoWord {
	var infoWords map[string]wiki.InfoWord

	file, err := os.Open(infoPath)
	if err != nil {
		log.Fatal(err)
		// TODO check if file is missing and just return empty
	}
	defer file.Close()

	// load the map

	return infoWords
}

func saveInfoWords(infoWords map[string]wiki.InfoWord) {
	json.Marshal(infoWords)
}
