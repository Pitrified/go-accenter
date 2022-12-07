package accenter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	rand "example.com/accenter/pkg/rand"
	wiki "example.com/accenter/pkg/wiki"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// A RecordHolder to deal with WikiRecords and related InfoWord at once.
//
// This might still be a part of the model.
type RecordHolder struct {
	iws map[wiki.Word]*InfoWord
	wrs map[wiki.Word]*wiki.WikiRecord

	db *gorm.DB
}

// Create a new RecordHolder.
func NewRecordHolder() *RecordHolder {

	// create the RecordHolder
	rh := &RecordHolder{}

	// build the path to the data folder
	dataFol, err := filepath.Abs(filepath.Join("..", "..", "dataset"))
	if err != nil {
		log.Fatal(err)
	}

	// find where the main database is
	dbPath := filepath.Join(dataFol, "accenter.db")
	// os.Remove(dbPath)

	// create a logger for the DB transactions
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Show ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	// create the database connection
	rh.db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		// failed to connect database
		log.Fatal(err)
	}

	// migrate the schema
	rh.db.AutoMigrate(&InfoWord{})

	// load the known InfoWords
	var iwsL []InfoWord
	result := rh.db.Find(&iwsL)
	fmt.Printf("We know %+v InfoWord.\n", result.RowsAffected)
	rh.iws = make(map[wiki.Word]*InfoWord)
	for _, iw := range iwsL {
		rh.iws[iw.Word] = &iw
	}

	// load the useful WikiRecords
	// find where the WikiRecords are
	wrPath := filepath.Join(dataFol, "wiki01.jsonl")
	// open the wr file
	wrFile, err := os.Open(wrPath)
	if err != nil {
		log.Fatal(err)
	}
	defer wrFile.Close()
	// will place useful WR here
	wikiRecords := map[wiki.Word]*wiki.WikiRecord{}
	// read the file line by line
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
	scanner := bufio.NewScanner(wrFile)
	for scanner.Scan() {
		// read the line and load the WikiRecord
		line := scanner.Text()
		var wr wiki.WikiRecord
		json.Unmarshal([]byte(line), &wr)
		// first we check if we know this word
		if _, ok := rh.iws[wr.Word]; !ok {
			// create a brand new one
			rh.iws[wr.Word] = NewInfoWord(wr.Word)
			// add this word to the database
			rh.db.Create(rh.iws[wr.Word])
		}
		// if it has accent and is not useless, we save it
		if rh.iws[wr.Word].HasAccent && !rh.iws[wr.Word].Useless {
			// fmt.Printf("Adding %s\n", result_struct.Word)
			wikiRecords[wr.Word] = &wr
		}
	}

	// TODO the weight of a useless or non-accented word will be 0
	// but we still waste time iterating over it
	// so it makes sense to only keep loaded the active words

	return rh
}

// Pick a random word according to the weights.
//
// PickWeighted still lives in rand.extract.
func (rh *RecordHolder) ExtractRandWord() wiki.Word {
	// the weights are maintained in iws
	// the interface of Pick is still to be defined
	return rand.Pick(rh.iws)
}

// Add an error to the requested word.
//
// MAYBE also remove the errors.
func (rh *RecordHolder) AddError(word wiki.Word) {
	rh.iws[word].Errors += 1
	// TODO write to database
	// TODO recompute weights
}

// Set the useless state of the word.
func (rh *RecordHolder) MarkUseless(word wiki.Word, useless bool) {
	rh.iws[word].Useless = useless
	// TODO write to database
	// TODO recompute weights
}

// we might avoid loading the info words
//
// just do a query every time
// while loading the records
//  if the word exists:
//     if it's not useless: load it
//     if it is useless   : skip it
//  if the word does not exists:
//     create it in the database
//
// how to compute the weight?
// max weight: one query sum(weights)
// compute running sum: https://stackoverflow.com/a/58339386/20222481
// generate random number and select the max where running is less than random
//
// or
//
// we load all the info
//
// if miss info while loading records: add it to the database (and the map)
// compute the weight (while loading) and the max weight
// generate rand [0, max weight] and range over the map, remove the weight
// when you change a info: change the map and the database, update the max weight

// // Load a dataset and attach a weight to each word.
// //
// // First the info, then the words.
// // Use the info to compute the weights on the fly,
// // no need for a weighted record.
// //
// // A function of the model so we have access to the InfoWords
// // Named LoadWikiRecords, more specific.
// func LoadDataset() map[wiki.Word]*wiki.WikiRecord {
//
// 	// wikiPath := findDataset("wikiRecords10")
// 	wikiPath := FindDataset("wikiRecords1k")
// 	wikiRecords := loadWikiRecords(wikiPath)
// 	fmt.Printf("Loaded %d WikiRecord\n", len(wikiRecords))
//
// 	infoPath := FindDataset("infoRecords")
// 	infoWords := loadInfoWords(infoPath)
// 	fmt.Printf("Loaded %d InfoWord\n", len(infoWords))
//
// 	// if we have some WikiRecord and no InfoWord for them create the default info
// 	// TODO update the iws and the db
// 	for word := range wikiRecords {
// 		if _, ok := infoWords[word]; !ok {
// 			infoWords[word] = &InfoWord{
// 				Word:      word,
// 				Errors:    0,
// 				Frequency: 1,
// 			}
// 		}
// 	}
//
// 	// FIXME should not load InfoWords for word we do not have in the WikiRecord
// 	// but then should also not delete them from the file?
// 	// almost as if we needed a database...
//
// 	// save the updated info
// 	SaveInfoWords(infoPath, infoWords)
//
// 	return wikiRecords
// }

// // Given a dataset identifier, get the absolute path to it.
// //
// // TODO a function of the model
// // or just in utils ?
// // it's kinda in between, the Abs(Join(baseFol, dbName)) is an util func
// // the case is a part of the model
// // might be two const at the top of this file
// // (as whichDataset := map[string]string)
// // no this is all useless just join the Abs(dataFol) and the file name
// func FindDataset(whichDataset string) string {
// 	dataFol := filepath.Join("..", "..", "dataset")
//
// 	dataName := ""
//
// 	switch whichDataset {
// 	case "wikiRecords10":
// 		dataName = "wiki01.jsonl"
// 	case "wikiRecords1k":
// 		dataName = "kaikki.org-dictionary-French-1k-accent.jsonl"
// 	case "infoRecords":
// 		dataName = "info01.json"
// 	case "infoRecordsDB":
// 		dataName = "info01.db"
// 	default:
// 		// this is fairly bad
// 		log.Fatalf("Unrecognized dataset tag to load: %s.\n", whichDataset)
// 	}
//
// 	dataPath := filepath.Join(dataFol, dataName)
// 	dataPath, err := filepath.Abs(dataPath)
// 	if err != nil {
// 		// if we cannot even build the path of the wiki data file,
// 		// we have a big problem
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Loading %s from %s\n", whichDataset, dataPath)
// 	return dataPath
// }

// // Load the wiki records.
// //
// // TODO a function of WikiRecords (actually NewWR() WRs)
// // If we want to avoid loading all records,
// // we should pass a map of acceptable words.
// // But if we want to also add InfoWord for an unknown word,
// // we need to
// // * load all records, then range over them and add the word
// // * have a reference to iws inside this func
// // * iws could be a broader WordData object, so we have all the attributes at once
// //
// // read a file line by line
// // https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
// func loadWikiRecords(wikiPath string) map[wiki.Word]*wiki.WikiRecord {
// 	file, err := os.Open(wikiPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	wikiRecords := map[wiki.Word]*wiki.WikiRecord{}
//
// 	scanner := bufio.NewScanner(file)
// 	// optionally, resize scanner's capacity for lines over 64K
// 	for scanner.Scan() {
// 		var result_struct wiki.WikiRecord
// 		line := scanner.Text()
// 		json.Unmarshal([]byte(line), &result_struct)
// 		if utils.IsAccentedWord(result_struct.Word) {
// 			// fmt.Printf("Adding %s\n", result_struct.Word)
// 			wikiRecords[result_struct.Word] = &result_struct
// 		}
// 	}
//
// 	// I mean is it really that bad?
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return wikiRecords
// }

// // Load the info we have for each word.
// // TODO a function of InfoWords
// func loadInfoWords(infoPath string) map[wiki.Word]*InfoWord {
// 	infoWords := map[wiki.Word]*InfoWord{}
//
// 	file, err := os.Open(infoPath)
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 		// if file is missing just return empty
// 		// the error might be something else
// 		return infoWords
// 	}
// 	defer file.Close()
//
// 	// load the map: read the jsonFile as a byte array
// 	byteValue, _ := io.ReadAll(file)
// 	json.Unmarshal([]byte(byteValue), &infoWords)
//
// 	return infoWords
// }

// // TODO a function of InfoWords
// func SaveInfoWords(infoPath string, infoWords map[wiki.Word]*InfoWord) {
//
// 	fmt.Printf("Saving %d InfoWord\n", len(infoWords))
// 	byteValue, err := json.MarshalIndent(infoWords, "", " ")
// 	if err != nil {
// 		// I mean it's kinda bad but not a total failure
// 		fmt.Printf("%s\n", err)
// 	}
// 	err = os.WriteFile(infoPath, byteValue, 0644)
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 	}
//
// }
