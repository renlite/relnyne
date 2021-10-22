# Declarative Style for [Fyne UI](https://github.com/fyne-io/fyne) 

This library is still in development and is not ready to be used.

Relnyne is a very lightweigt extention of [Fyne Widgets](https://github.com/fyne-io/fyne/tree/master/widget) to describe the UI in a declarative way and offers a simple data and state container to manage the state and update. 

```go
package main

import (
	rl "renlite/relnyne"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	w "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello")

	win.SetContent(container.NewVBox(
		rl.Lable("Fyne!", "#01"),
		w.NewButton("Say Hello!", clickHi),
	))

	win.ShowAndRun()
}

func clickHi(){
	// Option 1
	rl.GetLabel("#01").SetText("Hello Fyne!")
	// Option 2
	rl.GetWidget("#01").(*w.Label).SetText("Hello Fyne!")
	// Option 3? with Generics
	// rl.GetWidget[*w.Label]("#01")
}
```
