package window

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/desktop"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
)

func NewApp() fyne.App {
	return fyne.NewAppWithDriver(desktop.NewEFLDriver())
}

func ShowApp() {
	app := NewApp()

	top := widget.NewEntry()
	bottom := widget.NewEntry()
	left := widget.NewEntry()
	right := widget.NewEntry()
	middle := widget.NewLabel("BorderLayout")
	middle.Alignment = fyne.TextAlignCenter

	borderLayout := layout.NewBorderLayout(top, bottom, left, right)
	container := fyne.NewContainerWithLayout(borderLayout,
		top, bottom, left, right, middle)

	window := app.NewWindow("First Project")
	window.Fullscreen()
	window.Canvas().SetScale(5.5)
	window.SetContent(container)
	window.SetContent(widget.NewVBox(
	))
	fyne.NewPos(0, 300)
	window.SetContent(widget.NewVBox(
		widget.NewLabel("Hello again"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	window.Show()
}
