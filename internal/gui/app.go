package accenter

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type guiApp struct {
	c *guiController

	kb *keyboard

	word    *canvas.Text
	glosses *widget.Label

	hint *widget.Button
	next *widget.Button

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
	a.word = canvas.NewText("____ (4)", theme.ForegroundColor())
	a.word.TextSize = 30
	a.word.Alignment = fyne.TextAlignCenter
	a.word.TextStyle = fyne.TextStyle{Bold: false, Monospace: true}
	// meaning of the word
	a.glosses = widget.NewLabel("Glosses:")
	a.glosses.Wrapping = fyne.TextWrapWord
	// assemble word info
	contWord := container.NewVBox(
		layout.NewSpacer(),
		a.word,
		a.glosses,
		layout.NewSpacer(),
	)

	// elements in the bottom of the main screen
	a.hint = widget.NewButton("Hint", a.hintCB)
	a.next = widget.NewButton("Next", a.nextCB)
	contControl := container.NewGridWithColumns(2, a.hint, a.next)

	// build the main screen
	contMain := container.NewBorder(
		nil, contControl, nil, nil,
		contWord,
	)

	// ##### ASSEMBLE #####
	borderCont := container.NewBorder(nil, contKeyboard, nil, nil,
		container.NewPadded(contMain),
	)
	a.mainWin.SetContent(borderCont)

}

// -------------------------------------------------------------------
//  Reactions to user input:
//  callbacks to communicate with the Controller
// -------------------------------------------------------------------

// Clicked the button requesting a hint.
func (a *guiApp) hintCB() {
	a.c.clickedHint()
}

// Clicked the button requesting a next.
func (a *guiApp) nextCB() {
	a.c.clickedNext()
}

// -------------------------------------------------------------------
//  Update the app UI:
//  new state received from the controller
// -------------------------------------------------------------------

// Update the word to find with the current state.
func (a *guiApp) updateWord(word string) {
	a.word.Text = word
	a.word.Refresh()
}

// Update the glosses info.
func (a *guiApp) updateGlossesInfo(glosses string) {
	a.glosses.SetText(glosses)
}
