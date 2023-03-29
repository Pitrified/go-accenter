package utils

import (
	"strings"
	"unicode"

	"accenter/pkg/wiki"

	mapset "github.com/deckarep/golang-set/v2"
)

var AccentedLetters = "âàéèëêïîôœüùûç"
var AccentedLettersSet = mapset.NewSet([]rune(AccentedLetters)...)

var StandardLetters = "qwertyuiopasdfghjklzxcvbnm"
var AllLetters = AccentedLetters + StandardLetters
var AllLettersSet = mapset.NewSet([]rune(AllLetters)...)

var UnaccentLetterMap = map[rune]rune{
	'â': 'a', 'à': 'a',
	'é': 'e', 'è': 'e', 'ë': 'e', 'ê': 'e',
	'ï': 'i', 'î': 'i',
	'ô': 'o', 'œ': 'o',
	'ü': 'u', 'ù': 'u', 'û': 'u',
	'ç': 'c',
	'À': 'A', 'Â': 'A',
	'É': 'E', 'È': 'E', 'Ê': 'E', 'Ë': 'E',
	'Î': 'I', 'Ï': 'I',
	'Œ': 'O', 'Ô': 'O',
	'Ù': 'U', 'Û': 'U', 'Ü': 'U',
	'Ç': 'C',
}

// Return the unaccented version of the rune.
func UnaccentLetter(r rune) rune {
	ur := UnaccentLetterMap[r]
	if ur != 0 {
		return ur
	} else {
		return r
	}
}

// Return the unaccented version of the word.
func UnaccentWord(word wiki.Word) wiki.Word {
	return wiki.Word(strings.Map(UnaccentLetter, string(word)))
}

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

// Return true if any letter in the word is accented.
func IsAccentedWord(word wiki.Word) bool {
	for _, letter := range word {
		if IsAccentedLetter(letter) {
			return true
		}
	}
	return false
}
