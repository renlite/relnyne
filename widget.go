package renlite

import (
	"fyne.io/fyne/v2/widget"
)

// Textgrid
func TextGrid(id ...string) *widget.TextGrid {
	grid := widget.NewTextGrid()
	if id[0] != "" {
		getSrv().widgets[id[0]] = grid
	}
	return grid
}

func GetTextGrid(id string) *widget.TextGrid {
	return GetWidget(id).(*widget.TextGrid)
}

// Label
func Lable(text string, id ...string) *widget.Label {
	label := widget.NewLabel(text)
	if id[0] != "" {
		getSrv().widgets[id[0]] = label
	}
	return label
}

func GetLabel(id string) *widget.Label {
	return GetWidget(id).(*widget.Label)
}

// Button
func Button(label string, tapped func(), id ...string) *widget.Button {
	button := widget.NewButton("Hi Ui OpenGL!", tapped)
	if id[0] != "" {
		getSrv().widgets[id[0]] = button
	}
	return button
}

func GetButton(id string) *widget.Button {
	return GetWidget(id).(*widget.Button)
}

