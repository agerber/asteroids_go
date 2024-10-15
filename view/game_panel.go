package view

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type GamePanel struct {
	gameFrame *GameFrame
}

func NewGamePanel(width int, height int) *GamePanel {
	gameFrame := NewGameFrame("Game Base", width, height)
	return &GamePanel{gameFrame: gameFrame}
}

func (gamePanel *GamePanel) drawOneShip() {
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}
	gamePanel.gameFrame.Add(da)
}
