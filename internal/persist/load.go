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

func LoadWeiDataset() []wiki.WeiWikiRecord {

	dataPath := findDataset()
	wikiRecords := loadWikiRecords(dataPath)

	wei_wrecords := make([]wiki.WeiWikiRecord, len(wikiRecords))
	for i, wr := range wikiRecords {
		wei_wrecords[i] = wiki.NewWeiWikiRecord(wr, 1)
	}

	return wei_wrecords
}

func findDataset() string {
	dataPath := filepath.Join("..", "..", "dataset", "wiki01.jsonl")
	dataPath, err := filepath.Abs(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataPath)
	return dataPath
}

// read a file line by line
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
func loadWikiRecords(dataPath string) []wiki.WikiRecord {
	file, err := os.Open(dataPath)
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
