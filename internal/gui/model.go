package accenter

import (
	"fmt"
	"strings"
	"unicode"

	persist "example.com/accenter/internal/persist"
	utils "example.com/accenter/internal/utils"
	weightedrand "example.com/accenter/internal/weightedrand"
	wiki "example.com/accenter/pkg/wiki"
)

// --------------------------------------------------------------------------------
//  Define and create the model
// --------------------------------------------------------------------------------

// Which kind of prompt to show
type hintLevel int

const (
	hintOff     hintLevel = iota // show just a placeholder char
	hintLetters                  // show unaccented letters
	hintAll                      // show real letters
)

type guiModel struct {
	wr map[wiki.Word]*wiki.WikiRecord
	iw map[wiki.Word]*weightedrand.InfoWord

	secretWord  wiki.Word // secret word to write
	currentChar int       // index of the next letter to write
	showWord    string    // word to show to the user
	hintOn      hintLevel // which kind of prompt to show
	glossesInfo string    // glosses related to the word

	lastMistake rune // some signal to communicate with the controller
}

func newModel() *guiModel {
	// create the model
	m := &guiModel{}

	// load the records and the info
	m.wr, m.iw = persist.LoadDataset()

	// pick the first word to find
	m.pickNewSecretWord()

	return m
}

// Pick a word to find, update the relative info.
func (m *guiModel) pickNewSecretWord() {
	// fmt.Printf("M: Setting hintOn to false\n")
	// this need to happen before buildShowWord
	m.hintOn = hintOff

	m.secretWord = weightedrand.ExtractWord(m.iw)
	// m.secretWord = "no_raw_glossës"
	// m.secretWord = "Azraël"
	m.currentChar = 0
	m.buildShowWord()
	m.buildAllGlosses()
	m.lastMistake = ' '

	fmt.Printf("M: Picked %+v\n", m.secretWord)
	fmt.Printf("M: %+v\n", m.wr[m.secretWord].GetAllGlosses())
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
	// we only accept a char if it's correct, just copy that from the secret
	// fill with placeholders
	// add the len hint

	// always show the real secret word we already inserted
	m.showWord = m.secretWord.PrefixStr(m.currentChar)

	switch m.hintOn {
	case hintOff:
		// show just a placeholder char
		m.showWord += strings.Repeat("_", m.secretWord.Len()-m.currentChar)
	case hintLetters:
		// show unaccented letters
		suffix := m.secretWord.Suffix(m.currentChar)
		m.showWord += string(utils.UnaccentWord(suffix))
	case hintAll:
		// show real letters
		m.showWord += m.secretWord.SuffixStr(m.currentChar)
	}
	// add the len of the word
	m.showWord += fmt.Sprintf(" (%d)", m.secretWord.Len())

	// fmt.Printf("M: built %s\n", m.showWord)
}

// Build the definition of the word.
func (m *guiModel) buildAllGlosses() {
	allGlosses := m.wr[m.secretWord].GetAllGlosses()
	allSensesGlosses := make([]string, 5)
	for _, sense := range allGlosses {
		thisSenseGlosses := strings.Join(sense, "\n")
		allSensesGlosses = append(allSensesGlosses, thisSenseGlosses)
	}
	m.glossesInfo = strings.Join(allSensesGlosses, "\n")
}

// --------------------------------------------------------------------------------
//  React to user input: change the model state
// --------------------------------------------------------------------------------

// A new letter was inserted.
//
// If the letter is correct add it to the current word,
// skip next chars if they are not letters (spaces, hyphens).
// If it's wrong, mark a flag to disable wrong buttons
func (m *guiModel) clicked(letter rune) {
	// fmt.Printf("M: Clicked '%c'\n", letter)

	// get the next letter in the secret word
	nextSecretRune, _ := m.secretWord.RuneAt(m.currentChar)
	// convert it to lowercase
	nextSecretRune = unicode.ToLower(nextSecretRune)
	if letter == nextSecretRune || true {
		// if it is correct, append it
		m.currentChar += 1
		m.lastMistake = ' '
	} else {
		// mark the mistake and exit, no need to update the rest
		m.lastMistake = letter
		return
	}

	// eat the next chars if they are not letters
	for _, letter := range m.secretWord.Suffix(m.currentChar) {
		// fmt.Printf("letter '%c'\n", letter)
		if !utils.AllLettersSet.Contains(letter) {
			m.currentChar += 1
		} else {
			break
		}
	}

	if m.currentChar == m.secretWord.Len() {
		// fmt.Printf("M: You won!\n")
		m.lastMistake = '!'
	}

	m.buildShowWord()
}

// Clicked the button requesting a hint.
func (m *guiModel) clickedHint() {
	if m.hintOn < hintAll {
		m.hintOn += 1
	}
	m.buildShowWord()
}
