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
	wr map[wiki.Word]*wiki.WikiRecord
	iw map[wiki.Word]*InfoWord

	secretWord  wiki.Word // secret word to write
	currentChar int       // index of the next letter to write
	ShowWord    string    // word to show to the user
	showHint    hintLevel // which kind of prompt to show
	GlossesInfo string    // glosses related to the word

	LastMistake rune // some signal to communicate with the controller
}

func NewModel() *GuiModel {
	// create the model
	m := &GuiModel{}

	// load the records and the info
	m.wr, m.iw = LoadDataset()

	// pick the first word to find
	m.PickNewSecretWord()

	return m
}

// Pick a word to find, update the relative info.
func (m *GuiModel) PickNewSecretWord() {
	// fmt.Printf("M: Setting hintOn to false\n")
	// this need to happen before buildShowWord
	m.showHint = hintOff

	m.secretWord = ExtractWord(m.iw)
	// m.secretWord = "no_raw_glossës"
	// m.secretWord = "Azraël"
	m.currentChar = 0
	m.buildShowWord()
	m.buildAllGlosses()
	m.LastMistake = ' '

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
		m.ShowWord += string(utils.UnaccentWord(suffix))
	case hintAll:
		// show real letters
		m.ShowWord += m.secretWord.SuffixStr(m.currentChar)
	}
	// add the len of the word
	m.ShowWord += fmt.Sprintf(" (%d)", m.secretWord.Len())

	// fmt.Printf("M: built %s\n", m.ShowWord)
}

// Build the definition of the word.
func (m *GuiModel) buildAllGlosses() {
	allGlosses := m.wr[m.secretWord].GetAllGlosses()
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
	if letter == nextSecretRune || true {
		// if it is correct, append it
		m.currentChar += 1
		m.LastMistake = ' '
	} else {
		// mark the mistake and exit, no need to update the rest
		m.LastMistake = letter
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
	m.iw[m.secretWord].Useless = true
	infoPath := FindDataset("infoRecords")
	SaveInfoWords(infoPath, m.iw)
}

// TODO move model to internal/model/model.go
// and place persist/load.go in model/load.go (as funcs of guiModel)
// weighted_rand split as well:
// move extract.pick to pkg/rand
// extract.ExtractWord to model/info_word.go
// and extract.ComputeWeight to model/info_word.go