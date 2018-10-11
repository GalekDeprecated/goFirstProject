package main

import (
	"github.com/fyne-io/fyne/desktop"
	"github.com/fyne-io/fyne/widget"
)

func ShowApp() {
	app := desktop.NewApp()

	window := app.NewWindow("First Project")
	window.SetContent(widget.NewVBox(
		widget.NewLabel("Hello again"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	window.Show()
}
