package wiki

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The number of runes in a Word is correct.
func TestWordLen(t *testing.T) {
	cases := []struct {
		w  Word
		wl int
	}{
		{"", 0},
		{"a", 1},
		{"Azraël", 6},
	}
	for _, c := range cases {
		gotLen := c.w.Len()
		Equal(t, gotLen, c.wl, c)
	}
}

func TestWordPrefix(t *testing.T) {
	cases := []struct {
		w  Word
		p  Word
		pl int
	}{
		{"", "", 0},
		{"", "", 1},
		{"a", "a", 1},
		{"Azraël", "Az", 2},
		{"a", "a", -1},
		{"a", "a", 2},
		{"Azraël", "Azraël", 12},
	}
	for _, c := range cases {
		gotPref := c.w.Prefix(c.pl)
		Equal(t, c.p, gotPref, c)
	}
}

func Equal(t *testing.T, expected, actual, _case any) {
	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Failed case %+v, got %+v, expected %+v", _case, actual, expected),
	)
}

func TestWordRuneAt(t *testing.T) {
	errEmptyWord := errors.New("empty word cannot have runes in a specific location")
	errOutOfWord := errors.New("tried to get rune out of the word")
	cases := []struct {
		w   Word
		r   rune
		e   error
		pos int
	}{
		{"", ' ', errEmptyWord, 0},
		{"word", ' ', errOutOfWord, -1},
		{"word", ' ', errOutOfWord, 6},
		{"Azraël", 'A', nil, 0},
		{"Azraël", 'ë', nil, 4},
	}
	for _, c := range cases {
		gotRune, gotErr := c.w.RuneAt(c.pos)
		if gotErr == nil {
			Equal(t, c.r, gotRune, c)
		} else {
			Equal(t, c.e.Error(), gotErr.Error(), c)
		}
	}
}
