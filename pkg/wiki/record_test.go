package accenter

import (
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
