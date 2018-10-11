package main

import (
	"github.com/fyne-io/fyne/desktop"
	"github.com/fyne-io/fyne/widget"
)

func main() {
	app := desktop.NewApp()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewList(
		widget.NewLabel("Hello everyone!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.Show()
}
