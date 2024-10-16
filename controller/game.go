package controller

import (
	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gamePanel *view.GamePanel
}

func NewGame() *Game {
	gamePanel := view.NewGamePanel(config.DIM)
	return &Game{gamePanel: gamePanel}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gamePanel.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return config.DIM.Width, config.DIM.Height
}
