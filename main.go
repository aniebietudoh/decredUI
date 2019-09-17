package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("decred")
	w.Resize(fyne.NewSize(340, 600))
	w.SetContent(widget.NewVBox(
		widget.NewLabelWithStyle("decred", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
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
