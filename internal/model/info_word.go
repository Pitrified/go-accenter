package accenter

import (
	utils "example.com/accenter/internal/utils"
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
	Word      wiki.Word `gorm:"primarykey"` // the word
	Errors    int
	Frequency int
	Useless   bool
	TimesSeen int
	HasAccent bool
	Weight    int
}

func NewInfoWord(word wiki.Word) *InfoWord {
	iw := &InfoWord{
		Word:      word,
		Errors:    0,
		Frequency: 1,
		Useless:   false,
		TimesSeen: 0,
		HasAccent: utils.IsAccentedWord(word),
	}
	iw.updateWeight()
	return iw
}

// Update the InfoWord weight according to the current state.
//
// Return the delta weight.
func (iw *InfoWord) updateWeight() int {
	oldWeight := iw.Weight

	// never show useless words
	if iw.Useless {
		iw.Weight = 0
		return iw.Weight - oldWeight
	}

	// show unseen words very often
	if iw.TimesSeen == 0 {
		iw.Weight = 100
		return iw.Weight - oldWeight
	}

	// boost words with errors
	// as a words is seen more often, show it less
	// MAYBE do 10/(TimesSeen-Errors) to boost errors even more
	iw.Weight = int(10/iw.TimesSeen + (iw.Errors * 5))

	// has to be at least one
	if iw.Weight == 0 {
		iw.Weight = 1
	}

	return iw.Weight - oldWeight
}

// Add an error to the word.
//
// Return the delta weight.
func (iw *InfoWord) AddError() int {
	iw.Errors += 1
	return iw.updateWeight()
}

// Remove an error from the word.
//
// Return the delta weight.
func (iw *InfoWord) RemoveError() int {
	if iw.Errors > 0 {
		iw.Errors -= 1
	}
	return iw.updateWeight()
}

// Mark a word as useless.
//
// Return the delta weight.
func (iw *InfoWord) MarkUseless() int {
	iw.Useless = true
	return iw.updateWeight()
}

// The word was seen.
//
// Return the delta weight.
func (iw *InfoWord) Seen() int {
	iw.TimesSeen += 1
	return iw.updateWeight()
}
