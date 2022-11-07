package accenter

import (
	rand "example.com/accenter/pkg/rand"
	wiki "example.com/accenter/pkg/wiki"
)

// Information on the words.
//
// We weigh the records using:
// - number of error done on this word
// - frequency of a word, un-normalized
// - number or times a word was seen, heavily boost those at 0
// - uselessness of a word
type InfoWord struct {
	Word      wiki.Word `json:"w"`
	Errors    int       `json:"e"`
	Frequency int       `json:"f"`
	Useless   bool      `json:"u"`
	TimeSeen  int       `json:"s"`
}

// given a map of InfoWord
// pick one according to some logic
func ExtractWord(m map[wiki.Word]*InfoWord) wiki.Word {
	return rand.Pick(m)
}

// With facilities to read/write InfoWords.
type InfoWords struct {
	iws map[wiki.Word]*InfoWord
}

// PickWeighted still lives in rand.extract.
func (iws *InfoWords) ExtractWord() wiki.Word {
	return rand.Pick(iws.iws)
}

// Load the info words in a location.
func LoadInfoWords(path string) *InfoWords {
	return &InfoWords{}
}

// do not recompute everything, we know the old weight,
// so just update the total with the delta
// so we just need to call the `ComputeWeight` func once
// but remember that we might have suddenly useless words
// that simply will have 0 weight so we solve it

// Load a map with the useful InfoWords.
//
// Frankly should not be a method of a single IW.
// We make a InfoWords type,
// that will have this method and hold the map of IWs.
func (iw *InfoWord) Map() map[wiki.Word]InfoWord {
	iws := map[wiki.Word]InfoWord{}
	// query all the words
	// where useless is false
	return iws
}
