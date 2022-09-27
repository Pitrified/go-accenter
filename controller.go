package main

type guiController struct {
	a *guiApp
}

// Create a new controller, linked to the view and the model
func newController() *guiController {
	c := &guiController{}

	// create the view
	c.a = newApp(c)

	return c
}

// Create the UI and run the app.
func (c *guiController) run() {

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
