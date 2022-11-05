package accenter

import (
	"fmt"
	"math/rand"
	"time"
)

type guiController struct {
	a *guiApp
	m *guiModel
}

// Create a new controller, linked to the view and the model.
func NewController() *guiController {

	// initialize the random
	rand.Seed(time.Now().UnixNano())

	c := &guiController{}

	// create the view
	c.a = newApp(c)

	// initialize the model
	c.m = newModel()

	// create the UI, using placeholders everywhere
	c.a.buildUI()

	// update all the moving parts to match the current state:
	// the model has reasonable default values,
	// the view has only placeholders
	c.initAll()

	return c
}

// Run the app.
func (c *guiController) Run() {
	// run the app (will block)
	c.a.runApp()
}

func (c *guiController) initAll() {
	c.updateWord()
}

// --------------------------------------------------------------------------------
//  Reacts to events from UI (the view calls these funcs from the callbacks):
//  change the state of the model, then update the view.
// --------------------------------------------------------------------------------

// A keyboard button was clicked.
func (c *guiController) clicked(letter rune) {
	fmt.Printf("C: Clicked '%c'\n", letter)
	// update the model
	c.m.clicked(letter)
	// update all the pieces of the view
	c.updateWord()
}

// --------------------------------------------------------------------------------
//  The model has changed:
//  the controller knows which view elements must be updated.
// --------------------------------------------------------------------------------

// The word to show has changed.
func (c *guiController) updateWord() {
	// get the word to show from the model
	// place it in the view
	c.a.updateWord(c.m.showWord)
}
