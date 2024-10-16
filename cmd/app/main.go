package main

import (
	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.DIM.Width, config.DIM.Height)
	ebiten.SetWindowTitle(config.WINDOW_TITLE)
	if err := ebiten.RunGame(controller.NewGame()); err != nil {
		panic(err)
	}
}
