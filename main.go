package main

import (
	"image/color"
	"log"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/layout"
	"gioui.org/ui/measure"
	"gioui.org/ui/paint"
	"gioui.org/ui/text"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.WithWidth(ui.Dp(400)),
			app.WithHeight(ui.Dp(720)),
			app.WithTitle("Decred UI"),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	regular, err := sfnt.Parse(goregular.TTF)
	if err != nil {
		panic("failed to load font")
	}
	var cfg app.Config
	var faces measure.Faces
	grey := color.RGBA{17, 0, 0, 190}
	face := faces.For(regular, ui.Sp(30))
	message := "Welcome to \nDecred iOS Wallet"
	ops := new(ui.Ops)
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.UpdateEvent:
			cfg = e.Config
			faces.Reset(&cfg)
			cs := layout.RigidConstraints(e.Size)
			ops.Reset()
			var material ui.MacroOp
			material.Record(ops)
			paint.ColorOp{Color: grey}.Add(ops)
			material.Stop()
			text.Label{Material: material, Face: face, Alignment: text.Start, Text: message}.Layout(ops, cs)
			w.Update(ops)
		}
	}
}
