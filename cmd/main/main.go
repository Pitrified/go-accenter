package main

import (
	gui "example.com/accenter/internal/gui"
	persist "example.com/accenter/internal/persist"
)

func main() {
	theController := gui.NewController()
	theController.Run()

	dataPath := persist.FindDataset()
	persist.LoadWikiRecords(dataPath)
}
