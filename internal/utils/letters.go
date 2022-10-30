package accenter

import (
	"unicode"

	wiki "example.com/accenter/pkg/wiki"

	mapset "github.com/deckarep/golang-set/v2"
)

var AccentedLetters = createAccentedSet()

func createAccentedSet() mapset.Set[rune] {
	accentedLettersSet := mapset.NewSet[rune]()
	// accentedLettersSlice := []rune("âàéèëêïîôœüùûç")
	// for _, letter := range accentedLettersSlice {
	for _, letter := range "âàéèëêïîôœüùûç" {
		// fmt.Printf("\t%v %c\n", letter, letter)
		accentedLettersSet.Add(letter)
		// or we could call unicode.ToLower() when checking
		// accentedLettersSet.Add(unicode.ToUpper(letter))
	}
	return accentedLettersSet
}

func IsAccentedLetter(letter rune) bool {
	lowLetter := unicode.ToLower(letter)
	return AccentedLetters.Contains(lowLetter)
}

func IsAccentedWord(word wiki.Word) bool {
	for _, letter := range word {
		if IsAccentedLetter(letter) {
			return true
		}
	}
	return false
}
