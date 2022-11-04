package accenter

import (
	"fmt"
	"math/rand"
	"time"

	utils "example.com/accenter/internal/utils"
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

	fmt.Printf("AL %+v\n", utils.AccentedLetters)

	// create the view
	c.a = newApp(c)

	// initialize the model
	c.m = newModel()

	// create the UI, using placeholders everywhere
	c.a.buildUI()

	// update all the moving parts to match the current state:
	// the model has reasonable default values,
	// the view has only placeholders
	// c.initAll()

	return c
}

// Run the app.
func (c *guiController) Run() {
	// run the app (will block)
	c.a.runApp()
}
