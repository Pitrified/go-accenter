package accenter

import (
	wiki "example.com/accenter/pkg/wiki"
)

// // A WikiRecord with attached a weight, used to select which word to show.
// type WeiWikiRecord struct {
// 	WikiRecord WikiRecord
// 	Weight     int
// }

// func NewWeiWikiRecord(wr WikiRecord, weight int) WeiWikiRecord {
// 	return WeiWikiRecord{
// 		WikiRecord: wr,
// 		Weight:     weight,
// 	}
// }

// Information on the words.
//
// We weigh the records using:
// - the number of error done on this word
// - the frequency of a word, unnormalized
//
// TODO Skip a word if it's marked as useless
// TODO Save how many time a word was shown, heavily boost those at 0
type InfoWord struct {
	Word      wiki.Word `json:"w"`
	Errors    int       `json:"e"`
	Frequency int       `json:"f"`
	Useless   bool      `json:"u"`
}

// basically we are reinventing a database, the key is the word
