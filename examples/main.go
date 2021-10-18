package main

import (
	rl "renlite/go_fyne"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	w "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello")

	win.SetContent(container.NewVBox(
		rl.Lable("Fyne!", "#01"),
		w.NewButton("Say Hello!", click),
	))

	win.ShowAndRun()
}

func click(){
	// Option 1
	rl.GetLabel("#01").SetText("Hello Fyne!")
	// Option 2
	rl.GetWidget("#01").(*w.Label).SetText("Hello Fyne!")
	// Option 3? with Generics
	// rl.GetWidget[*w.Label]("#01")
}