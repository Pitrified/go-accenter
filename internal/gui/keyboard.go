package accenter

import (
	"fmt"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define a keyboard.
type keyboard struct {
	a *guiApp

	keys [][]*widget.Button
}

// Create a new keyboard.
func newKeyboard(a *guiApp) *keyboard {
	return &keyboard{a: a}
}

// Build the keyboard with a lot of buttons.
// àÀ âÂ éÉ èÈ êÊ ëË îÎ ïÏ œŒ ôÔ ùÙ ûÛ üÜ çÇ « » €
func (kb *keyboard) buildKeyboard() *fyne.Container {

	// the letters to show
	var rows [5][]rune
	rows[0] = []rune("âàéèëê")
	rows[1] = []rune("ïîôœüùûç")
	rows[2] = []rune("qwertyuiop")
	rows[3] = []rune("asdfghjkl")
	rows[4] = []rune(" zxcvbnm<")

	// prepare the rows of buttons
	kb.keys = make([][]*widget.Button, 5)

	// iterate over each row
	for i := 0; i < len(rows); i++ {
		row := rows[i]

		// prepare the button row
		kb.keys[i] = make([]*widget.Button, len(row))

		// build each button
		// fmt.Printf("%v %T %T %c\n", row, row, row[0], row)
		for ii := 0; ii < len(row); ii++ {

			// get the letter in the row
			letter := row[ii]
			// fmt.Printf("\t%v %c\n", letter, letter)

			// build the button
			kb.keys[i][ii] = widget.NewButton(
				fmt.Sprintf("%c", unicode.ToUpper(letter)),
				func() { kb.keysCB(letter) },
			)

		}
	}

	contKeys := container.NewVBox(
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					6,
					kb.keys[0][0], kb.keys[0][1], kb.keys[0][2], kb.keys[0][3], kb.keys[0][4],
					kb.keys[0][5],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					8,
					kb.keys[1][0], kb.keys[1][1], kb.keys[1][2], kb.keys[1][3], kb.keys[1][4],
					kb.keys[1][5], kb.keys[1][6], kb.keys[1][7],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					10,
					kb.keys[2][0], kb.keys[2][1], kb.keys[2][2], kb.keys[2][3], kb.keys[2][4],
					kb.keys[2][5], kb.keys[2][6], kb.keys[2][7], kb.keys[2][8], kb.keys[2][9],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					9,
					kb.keys[3][0], kb.keys[3][1], kb.keys[3][2], kb.keys[3][3], kb.keys[3][4],
					kb.keys[3][5], kb.keys[3][6], kb.keys[3][7], kb.keys[3][8],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					9,
					kb.keys[4][0], kb.keys[4][1], kb.keys[4][2], kb.keys[4][3], kb.keys[4][4],
					kb.keys[4][5], kb.keys[4][6], kb.keys[4][7], kb.keys[4][8],
				),
			),
		),
	)

	return contKeys
}

// -------------------------------------------------------------------
// Reactions to user input
// -------------------------------------------------------------------

// Clicked one of the keyoard buttons
func (kb *keyboard) keysCB(d rune) {
	fmt.Printf("Clicked '%c'\n", d)
	// kb.a.c.move(d)
}
