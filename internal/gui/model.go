package accenter

import (
	"fmt"
	"strings"
	"unicode/utf8"

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
	lastMistake   rune
}

func newModel() *guiModel {
	// create the model
	m := &guiModel{}

	// load the records and the info
	m.wr, m.iw = persist.LoadDataset()

	// pick a word to find
	m.secretWord = weightedrand.ExtractWord(m.iw)
	// m.secretWord = "Aborigène"
	m.secretWord = "Azraël"
	// m.secretWord = "no_raw_glossës"
	// m.secretWord = "Boquériny"
	m.secretWordLen = utf8.RuneCountInString(string(m.secretWord))
	m.currentWord = ""
	m.buildShowWord()
	m.buildAllGlosses()

	fmt.Printf("sad len %+v - rune count %+v\n",
		len(m.secretWord), utf8.RuneCountInString(string(m.secretWord)),
	)

	m.lastMistake = ' '

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
	// and we only accept a char if it's correct, so.
	// fill with placeholders
	// add the len hint

	// m.showWord = m.currentWord
	currentWordLen := utf8.RuneCountInString(m.currentWord)
	m.showWord = string([]rune(m.secretWord)[:currentWordLen])
	m.showWord += strings.Repeat("_", m.secretWordLen-currentWordLen)

	// m.showWord = strings.Repeat("_", m.secretWordLen)
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
	// fmt.Printf("M: Clicked '%c'\n", letter)

	m.currentWord += string(letter)
	fmt.Printf("m.currentWord %+v\n", m.currentWord)

	currentWordLen := utf8.RuneCountInString(m.currentWord)
	if currentWordLen == m.secretWordLen {
		fmt.Printf("You won!\n")
		// should communicate this somehow
	}

	m.buildShowWord()

	// if the letter is correct add it to the current word
	// skip next chars if they are not letters (spaces, hyphens)

	// if it's wrong, mark a flag -> will pop up a red button, possibly the keyboard one
	// we progressively disable buttons as the user clicks the wrong ones
}
