package main

func main() {
	// a := app.New()
	// w := a.NewWindow("Hello World")

	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.ShowAndRun()

	theController := newController()
	theController.run()
}
