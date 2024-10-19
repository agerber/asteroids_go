package main

import (
	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(common.DIM.Width, common.DIM.Height)
	ebiten.SetWindowTitle(common.WINDOW_TITLE)
	if err := ebiten.RunGame(controller.NewGame()); err != nil {
		panic(err)
	}
}
