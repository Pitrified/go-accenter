package main

import (
	"fmt"

	"accenter/internal/gui"
)

func main() {
	fmt.Println("Going.")

	theController := gui.NewController()
	theController.Run()
}
