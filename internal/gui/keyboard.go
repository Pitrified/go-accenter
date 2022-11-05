package accenter

import (
	"fmt"
	"unicode"

	utils "example.com/accenter/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define a keyboard.
type keyboard struct {
	a *guiApp

	keys map[rune]*widget.Button
}

// Create a new keyboard.
func newKeyboard(a *guiApp) *keyboard {
	return &keyboard{a: a}
}

// Build the keyboard with a lot of buttons.
// àÀ âÂ éÉ èÈ êÊ ëË îÎ ïÏ œŒ ôÔ ùÙ ûÛ üÜ çÇ « » €
func (kb *keyboard) buildKeyboard() *fyne.Container {

	// allLetters := "âàéèëêïîôœüùûçqwertyuiopasdfghjklzxcvbnm"
	kb.keys = make(map[rune]*widget.Button)

	for _, letter := range utils.AllLetters {
		letter := letter
		kb.keys[letter] = widget.NewButton(
			fmt.Sprintf("%c", unicode.ToUpper(letter)),
			func() { kb.keysCB(letter) },
		)
	}

	contKeys := container.NewVBox(
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					6,
					kb.keys['â'],
					kb.keys['à'],
					kb.keys['é'],
					kb.keys['è'],
					kb.keys['ë'],
					kb.keys['ê'],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					8,
					kb.keys['ï'],
					kb.keys['î'],
					kb.keys['ô'],
					kb.keys['œ'],
					kb.keys['ü'],
					kb.keys['ù'],
					kb.keys['û'],
					kb.keys['ç'],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					10,
					kb.keys['q'],
					kb.keys['w'],
					kb.keys['e'],
					kb.keys['r'],
					kb.keys['t'],
					kb.keys['y'],
					kb.keys['u'],
					kb.keys['i'],
					kb.keys['o'],
					kb.keys['p'],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					9,
					kb.keys['a'],
					kb.keys['s'],
					kb.keys['d'],
					kb.keys['f'],
					kb.keys['g'],
					kb.keys['h'],
					kb.keys['j'],
					kb.keys['k'],
					kb.keys['l'],
				),
			),
		),
		container.NewCenter(
			container.NewHBox(
				container.NewGridWithColumns(
					7,
					kb.keys['z'],
					kb.keys['x'],
					kb.keys['c'],
					kb.keys['v'],
					kb.keys['b'],
					kb.keys['n'],
					kb.keys['m'],
				),
			),
		),
	)

	return contKeys
}

// -------------------------------------------------------------------
//  Reactions to user input:
//  callbacks to communicate with the Controller
// -------------------------------------------------------------------

// Clicked one of the keyoard buttons.
func (kb *keyboard) keysCB(letter rune) {
	// fmt.Printf("K: Clicked '%c'\n", letter)
	kb.a.c.clicked(letter)
}

// -------------------------------------------------------------------
//  Update the app UI:
//  new state received from the controller
// -------------------------------------------------------------------

// Enable all the keyboard buttons.
func (kb *keyboard) enableAll() {
	for _, button := range kb.keys {
		button.Enable()
	}
}

// Disable all the keyboard buttons.
func (kb *keyboard) disableAll() {
	for _, button := range kb.keys {
		button.Disable()
	}
}

// Disable the requested keyboard button.
func (kb *keyboard) disable(letter rune) {
	kb.keys[letter].Disable()
}
