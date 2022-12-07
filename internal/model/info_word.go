package accenter

import (
	utils "example.com/accenter/internal/utils"
	wiki "example.com/accenter/pkg/wiki"

	_ "github.com/mattn/go-sqlite3"
)

// Information on the words.
//
// We weigh the records using:
// - number of error done on this word
// - frequency of a word, un-normalized
// - number or times a word was seen, heavily boost those at 0
// - uselessness of a word
type InfoWord struct {
	Word      wiki.Word `json:"w"` // the word
	Errors    int       `json:"e"`
	Frequency int       `json:"f"`
	Useless   bool      `json:"u"`
	TimesSeen int       `json:"s"`
	HasAccent bool      `json:"a"`
}

func NewInfoWord(word wiki.Word) *InfoWord {
	return &InfoWord{
		Word:      word,
		Errors:    0,
		Frequency: 1,
		Useless:   false,
		TimesSeen: 0,
		HasAccent: utils.IsAccentedWord(word),
	}
}

// // given a map of InfoWord
// // pick one according to some logic
// //
// // DEPRECATED: use InfoWords.ExtractWord()
// func ExtractWord(m map[wiki.Word]*InfoWord) wiki.Word {
// 	return rand.Pick(m)
// }

// // A collection of InfoWord.
// //
// // With facilities to read/write InfoWords.
// type InfoWords struct {
// 	iws map[wiki.Word]*InfoWord
// 	db  *sql.DB
// }
//
// // Load the info words in the given location.
// func NewInfoWords(pathDB string) *InfoWords {
//
// 	// create the info word holder
// 	iws := &InfoWords{}
//
// 	// open the database for the InfoWords
// 	db, err := sql.Open("sqlite3", pathDB)
// 	if err != nil {
// 		panic(err)
// 	}
// 	iws.db = db
//
// 	// create database table if not exists
// 	createStr := `CREATE TABLE IF NOT EXISTS infowords (
//         word TEXT PRIMARY KEY,
//         errors INTEGER,
//         frequency INTEGER,
//         useless BOOLEAN,
//         timesseen INTEGER
//     );`
// 	_, err = db.Exec(createStr)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// load all words
//
// 	return iws
// }

// // Pick a random word according to the weights.
// //
// // PickWeighted still lives in rand.extract.
// func (iws *InfoWords) ExtractRandWord() wiki.Word {
// 	// the weights are maintained in iws
// 	// the interface of Pick is still to be defined
// 	return rand.Pick(iws.iws)
// }

// // Add an error to the requested word.
// //
// // MAYBE also remove the errors.
// func (iws *InfoWords) AddError(word wiki.Word) {
// 	iws.iws[word].Errors += 1
// 	// TODO write to database
// 	// TODO recompute weights
// }

// // Set the useless state of the word.
// func (iws *InfoWords) MarkUseless(word wiki.Word, useless bool) {
// 	iws.iws[word].Useless = useless
// 	// TODO write to database
// 	// TODO recompute weights
// }

// do not recompute everything, we know the old weight,
// so just update the total with the delta
// so we just need to call the `ComputeWeight` func once
// but remember that we might have suddenly useless words
// that simply will have 0 weight so we solve it

// // Load a map with the useful InfoWords.
// //
// // We make a InfoWords type,
// // that will have this method and hold the map of IWs.
// //
// // DEPRECATED: use NewInfoWords
// func (iw *InfoWord) Map() map[wiki.Word]InfoWord {
// 	iws := map[wiki.Word]InfoWord{}
// 	// query all the words
// 	// where useless is false
// 	return iws
// }
