package renlite

import (
	"sync"

	"fyne.io/fyne/v2"
)

var once sync.Once

type service struct {
	widgets map[string]fyne.CanvasObject
	layouts map[string]fyne.Container
	data    map[string]interface{}
	status  map[string]string
}

var srv *service

func getSrv() *service {
	once.Do(func() { // <-- atomic, does not allow repeating
		srv = &service{
			widgets: make(map[string]fyne.CanvasObject),
			layouts: make(map[string]fyne.Container),
			data:    make(map[string]interface{}),
			status:  make(map[string]string),
		} // <-- thread safe
	})
	return srv
}

func GetWidget(id string) fyne.CanvasObject {
	return getSrv().widgets[id]
}

func GetLayout(id string) fyne.Container {
	return getSrv().layouts[id]
}

func GetStore(id string) interface{} {
	return getSrv().data[id]
}
