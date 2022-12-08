package accenter

import (
	"fmt"
	"strings"
	"unicode"

	utils "example.com/accenter/internal/utils"
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

type GuiModel struct {
	rh *RecordHolder

	secretWord  wiki.Word // secret word to write
	currentChar int       // index of the next letter to write
	ShowWord    string    // word to show to the user
	showHint    hintLevel // which kind of prompt to show
	GlossesInfo string    // glosses related to the word

	// TODO: might be a iota? with named states of what happened after the last click
	// we need to know which was the last letter but in the controller we
	// literally receive clicked(letter rune) so it's right there
	// also the state should kinda not live in the view, who knows
	LastMistake rune // some signal to communicate with the controller
	didMistake  bool // true when a mistake was made on the current word
}

func NewModel() *GuiModel {
	// create the model
	m := &GuiModel{}

	// load the records and the info
	m.rh = NewRecordHolder()

	// pick the first word to find
	m.PickNewSecretWord()

	return m
}

// --------------------------------------------------------------------------------
//  Update the model
// --------------------------------------------------------------------------------

// Pick a word to find, update the relative info.
func (m *GuiModel) PickNewSecretWord() {

	// this need to happen before buildShowWord
	m.showHint = hintOff

	// pick a random word
	m.secretWord = m.rh.ExtractRandWord()
	// m.secretWord = "no_raw_glossës"
	// m.secretWord = "Azraël"
	fmt.Printf("M: Picked %+v\n", m.secretWord)
	// fmt.Printf("M: %+v\n", m.rh.wrs[m.secretWord].GetAllGlosses())

	// reset the typing progress
	m.currentChar = 0
	m.buildShowWord()
	m.buildAllGlosses()
	m.LastMistake = ' '
	m.didMistake = false
}

// Build the prompt to show.
//
// Show the correctly inserted prefix, then some placeholders.
// Show the len of the word.
func (m *GuiModel) buildShowWord() {
	// write the correct matches
	// we only accept a char if it's correct, just copy that from the secret
	// fill with placeholders
	// add the len hint

	// always show the real secret word we already inserted
	m.ShowWord = m.secretWord.PrefixStr(m.currentChar)

	switch m.showHint {
	case hintOff:
		// show just a placeholder char
		m.ShowWord += strings.Repeat("_", m.secretWord.Len()-m.currentChar)
	case hintLetters:
		// show unaccented letters
		suffix := m.secretWord.Suffix(m.currentChar)
		m.ShowWord += "|" + string(utils.UnaccentWord(suffix))
	case hintAll:
		// show real letters
		m.ShowWord += "|" + m.secretWord.SuffixStr(m.currentChar)
	}
	// add the len of the word
	m.ShowWord += fmt.Sprintf(" (%d)", m.secretWord.Len())

	// fmt.Printf("M: built %s\n", m.ShowWord)
}

// Build the definition of the word.
func (m *GuiModel) buildAllGlosses() {

	if _, ok := m.rh.wrs[m.secretWord]; !ok {
		fmt.Printf("missing %+v in m.rh.wrs\n", m.secretWord)
		// will fail very soon
	}

	allGlosses := m.rh.wrs[m.secretWord].GetAllGlosses()
	allSensesGlosses := make([]string, 5)
	for _, sense := range allGlosses {
		thisSenseGlosses := strings.Join(sense, "\n")
		allSensesGlosses = append(allSensesGlosses, thisSenseGlosses)
	}
	m.GlossesInfo = strings.Join(allSensesGlosses, "\n")
}

// --------------------------------------------------------------------------------
//  React to user input: change the model state
// --------------------------------------------------------------------------------

// A new letter was inserted.
//
// If the letter is correct add it to the current word,
// skip next chars if they are not letters (spaces, hyphens).
// If it's wrong, mark a flag to disable wrong buttons
func (m *GuiModel) Clicked(letter rune) {
	// fmt.Printf("M: Clicked '%c'\n", letter)

	// get the next letter in the secret word
	nextSecretRune, _ := m.secretWord.RuneAt(m.currentChar)
	// convert it to lowercase
	nextSecretRune = unicode.ToLower(nextSecretRune)
	if letter == nextSecretRune {
		// if it is correct, append it
		m.currentChar += 1
		m.LastMistake = ' '
	} else {
		// mark the mistake and exit, no need to update the rest
		m.LastMistake = letter
		if !m.didMistake {
			m.rh.AddError(m.secretWord)
			m.didMistake = true
		}
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
		m.LastMistake = '!'
		if !m.didMistake {
			m.rh.RemoveError(m.secretWord)
		}
	}

	m.buildShowWord()
}

// Clicked the button requesting a hint.
func (m *GuiModel) ClickedHint() {
	if m.showHint < hintAll {
		m.showHint += 1
	}
	m.buildShowWord()
}

// Clicked the button to mark a word as useless.
func (m *GuiModel) ClickedUseless() {
	m.rh.MarkUseless(m.secretWord, true)
}
