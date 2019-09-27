package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("decred")
	w.Resize(fyne.NewSize(340, 600))

	text1 := canvas.NewText(" decred", color.RGBA{0x00, 0x00, 0xdd, 0xff})
	container := fyne.NewContainer(text1)

	w.SetContent(widget.NewVBox(
		container,
		//widget.NewLabelWithStyle("decred", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Welcome to \nDecred iOS Wallet", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewButton("Create new wallet", func() {
			a.Quit()
		}),
		widget.NewButton("Restore an existing wallet", func() {
			a.Quit()
		}),
	))
	w.ShowAndRun()
}
