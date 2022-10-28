package accenter

import (
	persist "example.com/accenter/internal/persist"
	weightedrand "example.com/accenter/internal/weightedrand"
	wiki "example.com/accenter/pkg/wiki"
)

type guiController struct {
	a  *guiApp
	wr map[wiki.Word]wiki.WikiRecord
	iw map[wiki.Word]weightedrand.InfoWord
}

// Create a new controller, linked to the view and the model
func NewController() *guiController {
	c := &guiController{}

	// load the model
	c.wr, c.iw = persist.LoadDataset()

	// fmt.Printf("%+v\n\n", c.wwr[0])

	// create the view
	c.a = newApp(c)

	return c
}

// Create the UI and run the app.
func (c *guiController) Run() {

	// initialize the model

	// create the UI, using placeholders everywhere
	c.a.buildUI()

	// update all the moving parts to match the current state:
	// the model has reasonable default values,
	// the view has only placeholders
	// c.initAll()

	// run the app (will block)
	c.a.runApp()

}
