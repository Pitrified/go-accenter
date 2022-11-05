package accenter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	utils "example.com/accenter/internal/utils"
	wiki "example.com/accenter/pkg/wiki"
)

// Load a dataset and attach a weight to each word.
//
// First the info, then the words.
// Use the info to compute the weights on the fly,
// no need for a weighted record.
func LoadDataset() (
	map[wiki.Word]*wiki.WikiRecord,
	map[wiki.Word]*InfoWord,
) {

	// wikiPath := findDataset("wikiRecords10")
	wikiPath := FindDataset("wikiRecords1k")
	wikiRecords := loadWikiRecords(wikiPath)
	fmt.Printf("Loaded %d WikiRecord\n", len(wikiRecords))

	infoPath := FindDataset("infoRecords")
	infoWords := loadInfoWords(infoPath)
	fmt.Printf("Loaded %d InfoWord\n", len(infoWords))

	// if we have some WikiRecord and no InfoWord for them create the default info
	for word := range wikiRecords {
		if _, ok := infoWords[word]; !ok {
			infoWords[word] = &InfoWord{
				Word:      word,
				Errors:    0,
				Frequency: 1,
			}
		}
	}

	// FIXME should not load InfoWords for word we do not have in the WikiRecord
	// but then should also not delete them from the file?
	// almost as if we needed a database...

	// save the updated info
	SaveInfoWords(infoPath, infoWords)

	return wikiRecords, infoWords
}

func FindDataset(whichDataset string) string {
	dataFol := filepath.Join("..", "..", "dataset")

	dataName := ""

	switch whichDataset {
	case "wikiRecords10":
		dataName = "wiki01.jsonl"
	case "wikiRecords1k":
		dataName = "kaikki.org-dictionary-French-1k-accent.jsonl"
	case "infoRecords":
		dataName = "info01.json"
	default:
		// this is fairly bad
		log.Fatalf("Unrecognized dataset tag to load: %s.\n", whichDataset)
	}

	dataPath := filepath.Join(dataFol, dataName)
	dataPath, err := filepath.Abs(dataPath)
	if err != nil {
		// if we cannot even build the path of the wiki data file,
		// we have a big problem
		log.Fatal(err)
	}
	fmt.Printf("Loading %s from %s\n", whichDataset, dataPath)
	return dataPath
}

// Load the wiki records.
//
// read a file line by line
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
func loadWikiRecords(wikiPath string) map[wiki.Word]*wiki.WikiRecord {
	file, err := os.Open(wikiPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wikiRecords := map[wiki.Word]*wiki.WikiRecord{}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		var result_struct wiki.WikiRecord
		line := scanner.Text()
		json.Unmarshal([]byte(line), &result_struct)
		if utils.IsAccentedWord(result_struct.Word) {
			// fmt.Printf("Adding %s\n", result_struct.Word)
			wikiRecords[result_struct.Word] = &result_struct
		}
	}

	// I mean is it really that bad?
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wikiRecords
}

// Load the info we have for each word.
func loadInfoWords(infoPath string) map[wiki.Word]*InfoWord {
	infoWords := map[wiki.Word]*InfoWord{}

	file, err := os.Open(infoPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		// if file is missing just return empty
		// the error might be something else
		return infoWords
	}
	defer file.Close()

	// load the map: read the jsonFile as a byte array
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal([]byte(byteValue), &infoWords)

	return infoWords
}

func SaveInfoWords(infoPath string, infoWords map[wiki.Word]*InfoWord) {

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
