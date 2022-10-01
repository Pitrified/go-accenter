package accenter

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type guiApp struct {
	c *guiController

	kb *keyboard

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
	contMain := widget.NewLabel("Mock")

	// ##### ASSEMBLE #####
	borderCont := container.NewBorder(nil, contKeyboard, nil, nil,
		container.NewCenter(contMain),
	)
	a.mainWin.SetContent(borderCont)

}
