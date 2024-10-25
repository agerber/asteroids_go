package main

import (
	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	defer common.CloseSound()
	common.SetCommandCenterInstance(controller.NewCommandCenter())
	ebiten.SetWindowSize(common.DIM.Width, common.DIM.Height)
	ebiten.SetWindowTitle(common.WINDOW_TITLE)
	ebiten.SetTPS(common.GOLANG_FRAMES_PER_SECOND)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetScreenClearedEveryFrame(false)
	if err := ebiten.RunGame(controller.NewGame()); err != nil {
		panic(err)
	}
}
