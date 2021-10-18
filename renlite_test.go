package renlite

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	createUi(test.NewApp())

	lable := GetLabel("#01")
	assert.Equal(t, "Fyne!", lable.Text)
	test.Tap(GetButton("#02"))
	assert.Equal(t, "Hello Fyne!", lable.Text)
}

func createUi(app fyne.App) {
	win := app.NewWindow("Hello")
	win.SetContent(container.NewVBox(
		Lable("Fyne!", "#01"),
		Button("Say Hello", click, "#02"),
	))
}

func click() {
	GetLabel("#01").SetText("Hello Fyne!")
}
