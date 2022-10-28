package accenter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	weightedrand "example.com/accenter/internal/weightedrand"
	wiki "example.com/accenter/pkg/wiki"
)

// Load a dataset and attach a weight to each word.
//
// First the info, then the words.
// Use the info to compute the weights on the fly,
// no need for a weighted record.
func LoadDataset() (
	map[wiki.Word]wiki.WikiRecord,
	map[wiki.Word]weightedrand.InfoWord,
) {

	wikiPath := findDataset("wikiRecords")
	wikiRecords := loadWikiRecords(wikiPath)
	fmt.Printf("Loaded %d WikiRecord\n", len(wikiRecords))

	infoPath := findDataset("infoRecords")
	infoWords := loadInfoWords(infoPath)
	fmt.Printf("Loaded %d InfoWord\n", len(infoWords))

	// if we have some WikiRecord and no InfoWord for them create the default info
	for word := range wikiRecords {
		if _, ok := infoWords[word]; ok {
			infoWords[word] = weightedrand.InfoWord{
				Word:      word,
				Errors:    0,
				Frequency: 1,
			}
		}
	}

	// save the updated info
	saveInfoWords(infoPath, infoWords)

	return wikiRecords, infoWords
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
func loadWikiRecords(wikiPath string) map[wiki.Word]wiki.WikiRecord {
	file, err := os.Open(wikiPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wikiRecords := map[wiki.Word]wiki.WikiRecord{}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	var result_struct wiki.WikiRecord
	for scanner.Scan() {
		line := scanner.Text()
		json.Unmarshal([]byte(line), &result_struct)
		wikiRecords[result_struct.Word] = result_struct
	}

	// I mean is it really that bad?
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wikiRecords
}

// Load the info we have for each word.
func loadInfoWords(infoPath string) map[wiki.Word]weightedrand.InfoWord {
	infoWords := map[wiki.Word]weightedrand.InfoWord{}

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

func saveInfoWords(infoPath string, infoWords map[wiki.Word]weightedrand.InfoWord) {

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
