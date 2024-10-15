package view

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

type GameFrame struct {
	*gtk.Window
}

func NewGameFrame(title string, width int, height int) *GameFrame {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(fmt.Sprintf("Unable to create window: %v", err))
	}

	window.SetTitle(title)
	window.SetSizeRequest(width, height)
	window.SetResizable(false)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		panic(fmt.Sprintf("Unable to create box: %v", err))
	}
	window.Add(box)

	window.Connect("destroy", func() {
		gtk.MainQuit()
		os.Exit(0)
	})

	window.ShowAll()

	return &GameFrame{Window: window}
}
