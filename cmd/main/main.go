package main

import (
	"fmt"

	gui "example.com/accenter/internal/gui"
)

func main() {
	fmt.Println("Going.")

	theController := gui.NewController()
	theController.Run()
}
