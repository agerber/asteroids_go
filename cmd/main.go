package main

import (
	"log"

	"github.com/agerber/asteroids_go/cmd/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gameController := controller.NewGameController()

	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Game Base")
	if err := ebiten.RunGame(gameController); err != nil {
		log.Fatal(err)
	}
}
