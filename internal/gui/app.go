package accenter

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type guiApp struct {
	c *guiController

	kb *keyboard

	word    *widget.Label
	glosses *widget.Label

	fyneApp fyne.App
	mainWin fyne.Window
}

func newApp(c *guiController) *guiApp {

	// create the fyne app
	fyneApp := app.NewWithID("com.pitrified.accenter")
	mainWin := fyneApp.NewWindow("Accenter")

	// create the accenter app
	theApp := &guiApp{fyneApp: fyneApp, mainWin: mainWin, c: c}

	// add the link for typed runes
	theApp.mainWin.Canvas().SetOnTypedKey(theApp.typedKey)

	return theApp
}

func (a *guiApp) runApp() {
	a.mainWin.Resize(fyne.NewSize(100, 600))
	a.mainWin.Show()
	a.fyneApp.Run()
}

func (a *guiApp) typedKey(ev *fyne.KeyEvent) {
	fmt.Printf("typedKey  = %+v %T\n", ev, ev)
	switch ev.Name {
	case fyne.KeyEscape:
		a.fyneApp.Quit()
	// case fyne.KeyW, fyne.KeyS, fyne.KeyD, fyne.KeyE, fyne.KeyA, fyne.KeyQ:
	// 	a.control(ev.Name)
	// case fyne.KeySpace:
	// 	a.c.togglePenState()
	// case fyne.KeyH:
	// 	a.s.miscHelpCB()
	// case fyne.KeyF, fyne.KeyF11:
	// 	a.toggleFullscreen()
	default:
	}
}

// -------------------------------------------------------------------
//  Build the app UI
// -------------------------------------------------------------------

func (a *guiApp) buildUI() {

	// ##### KEYBOARD #####
	a.kb = newKeyboard(a)
	contKeyboard := a.kb.buildKeyboard()

	// ##### MAIN SCREEN #####
	// word to find
	// a.word = widget.NewLabel("Word: .... ... ... .. .. . . (4)")
	a.word = widget.NewLabelWithStyle(
		"____ (4) - you actually never see this",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Bold:      false,
			Monospace: true,
		},
	)
	// a.word.Alignment = fyne.TextAlignCenter
	// meaning of the word
	a.glosses = widget.NewLabel("Glosses:\none\ntwo\na very long gloss that explains a lot of information about the word")
	a.glosses.Wrapping = fyne.TextWrapWord
	// assemble word info
	contWord := container.NewVBox(
		layout.NewSpacer(),
		a.word,
		a.glosses,
		layout.NewSpacer(),
	)
	// add more elements to the screen
	labelControl := widget.NewLabel("Mock for buttons: hint, next...")
	labelControl.Alignment = fyne.TextAlignCenter
	// build the main screen
	contMain := container.NewBorder(
		nil, labelControl, nil, nil,
		contWord,
	)

	// ##### ASSEMBLE #####
	// frankly this title is useless, meh
	labelTitle := widget.NewLabelWithStyle(
		"Accenter", fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
	)
	borderCont := container.NewBorder(labelTitle, contKeyboard, nil, nil,
		// container.NewCenter(contMain),
		container.NewPadded(contMain),
	)
	a.mainWin.SetContent(borderCont)

}

// -------------------------------------------------------------------
//  Update the app UI
// -------------------------------------------------------------------

// Update the word to find with the current state.
func (a *guiApp) updateWord(word string) {
	a.word.SetText(word)
}

// Update the glosses info.
func (a *guiApp) updateGlossesInfo(glosses string) {
	a.glosses.SetText(glosses)
}
