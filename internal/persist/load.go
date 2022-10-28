package accenter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	wiki "example.com/accenter/internal/wiki"
)

// Load a dataset and attach a weight to each word.
func LoadWeiDataset() []wiki.WeiWikiRecord {

	wikiPath := findDataset("wikiRecords")
	wikiRecords := loadWikiRecords(wikiPath)
	fmt.Printf("Loaded %d WikiRecord\n", len(wikiRecords))

	infoPath := findDataset("infoRecords")
	infoWords := loadInfoWords(infoPath)
	fmt.Printf("Loaded %d InfoWord\n", len(infoWords))

	wei_wrecords := make([]wiki.WeiWikiRecord, len(wikiRecords))
	for i, wr := range wikiRecords {

		// fmt.Printf("Doing %s '%s' %d\n",
		// 	wr.Word, infoWords[wr.Word].Word,
		// 	len(infoWords[wr.Word].Word),
		// )

		// if the word is not in the map, an empty string is returned
		if len(infoWords[wr.Word].Word) > 0 {
			// TODO compute the actual weight from errors and frequency
			wei_wrecords[i] = wiki.NewWeiWikiRecord(wr, 1)

		} else {
			// assign a default weight to this record
			wei_wrecords[i] = wiki.NewWeiWikiRecord(wr, 1)
			// create the info about this word
			infoWords[wr.Word] = wiki.InfoWord{
				Word:      wr.Word,
				Errors:    0,
				Frequency: 1,
			}
		}
	}

	saveInfoWords(infoPath, infoWords)

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

// Load the wiki records.
//
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
	// optionally, resize scanner's capacity for lines over 64K
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

// Load the info we have for each word.
func loadInfoWords(infoPath string) map[string]wiki.InfoWord {
	infoWords := map[string]wiki.InfoWord{}

	file, err := os.Open(infoPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		// TODO check if file is missing and just return empty
		// the error might be something else
		return infoWords
	}
	defer file.Close()

	// load the map: read the jsonFile as a byte array
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal([]byte(byteValue), &infoWords)

	return infoWords
}

func saveInfoWords(infoPath string, infoWords map[string]wiki.InfoWord) {

	fmt.Printf("Saving %d InfoWord\n", len(infoWords))
	byteValue, err := json.MarshalIndent(infoWords, "", " ")
	if err != nil {
		// I mean it's kinda bad but not a total failure
		fmt.Printf("%s\n", err)
	}
	err = os.WriteFile(infoPath, byteValue, 0644)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

}
