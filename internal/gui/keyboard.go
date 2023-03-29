package gui

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

	keys   [][]*widget.Button
	layout [][]rune
}

// Create a new keyboard.
func newKeyboard(a *guiApp) *keyboard {
	return &keyboard{a: a}
}

// Build the keyboard with a lot of buttons.
// àÀ âÂ éÉ èÈ êÊ ëË îÎ ïÏ œŒ ôÔ ùÙ ûÛ üÜ çÇ « » €
func (kb *keyboard) buildKeyboard() *fyne.Container {

	// TODO: sanity check that all letters are present
	kb.layout = [][]rune{
		{'â', 'à', 'é', 'è', 'ë', 'ê'},
		{'ï', 'î', 'ô', 'œ', 'ü', 'ù', 'û', 'ç'},
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}

	// create the buttons
	kb.keys = make([][]*widget.Button, len(kb.layout))
	for i := range kb.layout {
		kb.keys[i] = make([]*widget.Button, len(kb.layout[i]))
		for j := range kb.layout[i] {
			letter := kb.layout[i][j]
			kb.keys[i][j] = widget.NewButton(
				fmt.Sprintf("%c", unicode.ToUpper(letter)),
				func() { kb.keysCB(letter) },
			)
		}
	}

	// create the keyboard
	rows := make([]fyne.CanvasObject, len(kb.keys))
	for i := range kb.keys {
		rows[i] = container.NewCenter(
			container.NewGridWithColumns(
				len(kb.keys[i]),
				ToCO(kb.keys[i]...)...,
			),
		)
	}
	contKeys := container.NewVBox(rows...)

	return contKeys
}

// -------------------------------------------------------------------
//  Reactions to user input:
//  callbacks to communicate with the Controller
// -------------------------------------------------------------------

// Clicked one of the keyboard buttons.
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
	for i := range kb.keys {
		for j := range kb.keys[i] {
			button := kb.keys[i][j]
			button.Enable()
		}
	}
}

// Disable all the keyboard buttons.
func (kb *keyboard) disableAll() {
	for i := range kb.keys {
		for j := range kb.keys[i] {
			button := kb.keys[i][j]
			button.Disable()
		}
	}
}

// Disable the requested keyboard button.
func (kb *keyboard) disable(letter rune) {
	for i := range kb.layout {
		for j := range kb.layout[i] {
			if kb.layout[i][j] == letter {
				kb.keys[i][j].Disable()
				return
			}
		}
	}
}
