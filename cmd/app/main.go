package main

import (
	"os"

	"github.com/agerber/asteroids_go/controller"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(&os.Args)
	_ = controller.NewGame()
	gtk.Main()
}
