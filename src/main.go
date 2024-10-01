package main

import (
	"fmt"

	"github.com/agerber/asteroids_go/src/mvc/view"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("Hello, World!")
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("My Game")
	gameFrame := &view.GamePanel{}
	if err := ebiten.RunGame(gameFrame); err != nil {
		fmt.Println(err)
	}
}
