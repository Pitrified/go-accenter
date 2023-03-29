package gui

import "fyne.io/fyne/v2"

// Convert a slice of a more specific type to a slice of fyne.CanvasObject.
//
// Mildly inefficient, but it works.
func ToCO[T fyne.CanvasObject](objects ...T) []fyne.CanvasObject {
	objCO := make([]fyne.CanvasObject, len(objects))
	for i := range objCO {
		objCO[i] = objects[i]
	}
	return objCO
}
