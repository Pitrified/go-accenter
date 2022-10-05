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

// read a file line by line
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
func LoadWikiRecords(dataPath string) []wiki.WikiRecord {
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var records []wiki.WikiRecord

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		var result map[string]interface{}
		json.Unmarshal([]byte(line), &result)

		var result_struct wiki.WikiRecord
		json.Unmarshal([]byte(line), &result_struct)

		records = append(records, result_struct)

		fmt.Println(line)
		// fmt.Println(result)
		fmt.Printf("%+v\n\n", result_struct)

		// break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return records
}

func FindDataset() string {
	dataPath := filepath.Join("..", "..", "dataset", "wiki01.jsonl")
	dataPath, err := filepath.Abs(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataPath)
	return dataPath
}
