package accenter

import (
	"fmt"
	"strings"

	persist "example.com/accenter/internal/persist"
	weightedrand "example.com/accenter/internal/weightedrand"
	wiki "example.com/accenter/pkg/wiki"
)

type guiModel struct {
	wr map[wiki.Word]*wiki.WikiRecord
	iw map[wiki.Word]*weightedrand.InfoWord

	secretWord    wiki.Word
	secretWordLen int
	currentWord   string
	showWord      string
	glossesInfo   string
}

func newModel() *guiModel {
	// create the model
	m := &guiModel{}

	// load the records and the info
	m.wr, m.iw = persist.LoadDataset()

	// pick a word to find
	m.secretWord = weightedrand.ExtractWord(m.iw)
	// m.secretWord = "Aborigène"
	// m.secretWord = "Azraël"
	// m.secretWord = "no_raw_glossës"
	// m.secretWord = "Boquériny"
	m.secretWordLen = len(m.secretWord)
	m.currentWord = ""
	m.buildShowWord()
	m.buildAllGlosses()

	fmt.Printf("Picked %+v\n", m.secretWord)
	// fmt.Printf("%+v\n", m.wr[m.secretWord].Senses[0].RawGlosses)
	fmt.Printf("%+v\n", m.wr[m.secretWord].GetAllGlosses())

	return m
}

// --------------------------------------------------------------------------------
//  Update the model
// --------------------------------------------------------------------------------

// Build the prompt to show.
//
// Show the correctly inserted prefix, then some placeholders.
// Show the len of the word.
func (m *guiModel) buildShowWord() {
	// write the correct matches
	// fill with placeholders
	// add the len hint
	m.showWord = strings.Repeat("_", m.secretWordLen)
	m.showWord += fmt.Sprintf(" (%d)", m.secretWordLen)
	fmt.Printf("M: built %s\n", m.showWord)
}

// Build the definition of the word.
func (m *guiModel) buildAllGlosses() {
	allGlosses := m.wr[m.secretWord].GetAllGlosses()
	allSensesGlosses := make([]string, 5)
	for _, sense := range allGlosses {
		thisSenseGlosses := strings.Join(sense, "\n")
		allSensesGlosses = append(allSensesGlosses, thisSenseGlosses)
	}
	m.glossesInfo += strings.Join(allSensesGlosses, "\n")
}

// --------------------------------------------------------------------------------
//  React to user input: change the model state
// --------------------------------------------------------------------------------

// A new letter was inserted.
func (m *guiModel) clicked(letter rune) {
	fmt.Printf("M: Clicked '%c'\n", letter)
}
