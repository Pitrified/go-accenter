package accenter

import (
	"unicode"

	wiki "example.com/accenter/pkg/wiki"

	mapset "github.com/deckarep/golang-set/v2"
)

var AccentedLetters = "âàéèëêïîôœüùûç"
var AccentedLettersSet = mapset.NewSet([]rune(AccentedLetters)...)

var StandardLetters = "qwertyuiopasdfghjklzxcvbnm"
var AllLetters = AccentedLetters + StandardLetters

// Return true if letter is accented.
//
// or we could convert the string to lower,
// turn it into a set,
// compute the intersection,
// check the len.
func IsAccentedLetter(letter rune) bool {
	lowLetter := unicode.ToLower(letter)
	return AccentedLettersSet.Contains(lowLetter)
}

func IsAccentedWord(word wiki.Word) bool {
	for _, letter := range word {
		if IsAccentedLetter(letter) {
			return true
		}
	}
	return false
}
