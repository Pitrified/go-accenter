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

	totalWeight int // the sum of the weights of the InfoWords

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
		// copy the word in a new location
		iwc := iw
		rh.iws[iw.Word] = &iwc
	}

	// load the useful WikiRecords
	// find where the WikiRecords are
	// wrFileName := "wiki01.jsonl"
	wrFileName := "kaikki.org-dictionary-French-1k-accent.jsonl"
	wrPath := filepath.Join(dataFol, wrFileName)
	// open the wr file
	wrFile, err := os.Open(wrPath)
	if err != nil {
		log.Fatal(err)
	}
	defer wrFile.Close()
	// will place useful WR here
	rh.wrs = map[wiki.Word]*wiki.WikiRecord{}
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
		// if rh.iws[wr.Word].HasAccent && !rh.iws[wr.Word].Useless {
		if rh.IsInteresting(wr.Word) {
			// fmt.Printf("Adding %s\n", wr.Word)
			rh.wrs[wr.Word] = &wr
		}
	}

	log.Printf("Loaded %d WikiRecords", len(rh.wrs))

	// the weight of a non-interesting word will be 0 but we still waste time
	// iterating over it so only keep the active words
	for _, iw := range rh.iws {
		// fmt.Printf("have %s\n", iw.Word)
		if !rh.IsInteresting(iw.Word) {
			delete(rh.iws, iw.Word)
			fmt.Printf("Deleting %s\n", iw.Word)
			continue
		}
		// fmt.Printf("Adding weight %s\n", wr.Word)
		rh.totalWeight += iw.Weight
	}
	log.Printf("InfoWord total weight %d", rh.totalWeight)

	return rh
}

// Decide if an InfoWord is interesting.
func (rh *RecordHolder) IsInteresting(word wiki.Word) bool {
	return (rh.iws[word].HasAccent && !rh.iws[word].Useless)
}

// Pick a random word according to the weights.
//
// PickWeighted still lives in rand.extract.
func (rh *RecordHolder) ExtractRandWord() wiki.Word {
	// the weights are maintained in iws
	// the interface of Pick is still to be defined

	// return rand.Pick(rh.iws)

	// pick the word
	word := rand.PickMap(
		rh.iws,
		func(iw *InfoWord) int { return iw.Weight },
		rh.totalWeight,
	)

	// mark as seen, update weight, write to DB
	rh.totalWeight += rh.iws[word].Seen()
	rh.UpdateWord(word)

	return word
}

// Add an error to the requested word.
func (rh *RecordHolder) AddError(word wiki.Word) {
	// track the new total weight
	rh.totalWeight += rh.iws[word].AddError()
	// write to database
	rh.UpdateWord(word)
}

// Remove an error from the requested word.
func (rh *RecordHolder) RemoveError(word wiki.Word) {
	// track the new total weight
	rh.totalWeight += rh.iws[word].RemoveError()
	// write to database
	rh.UpdateWord(word)
}

// Set the useless state of the word.
func (rh *RecordHolder) MarkUseless(word wiki.Word, useless bool) {
	// track the new total weight
	rh.totalWeight += rh.iws[word].MarkUseless()
	// write to database
	rh.UpdateWord(word)
}

// Update a word info in the database.
func (rh *RecordHolder) UpdateWord(word wiki.Word) {

	// https://gorm.io/docs/update.html#Save-All-Fields
	rh.db.Save(*rh.iws[word])

	// https://gorm.io/docs/update.html#Update-Selected-Fields
	// we could use this to specifically update only the changed fields
	// but seems like an hassle

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
//
// do not recompute everything, we know the old weight,
// so just update the total with the delta
// so we just need to call the `ComputeWeight` func once
// but remember that we might have suddenly useless words
// that simply will have 0 weight so we solve it
