package controller

import "github.com/agerber/asteroids_go/view"

const (
	GAME_SCREEN_WIDTH  = 1500
	GAME_SCREEN_HEIGHT = 950
)

type Game struct {
	gamePanel *view.GamePanel
}

func NewGame() *Game {
	gamePanel := view.NewGamePanel(GAME_SCREEN_WIDTH, GAME_SCREEN_HEIGHT)
	return &Game{gamePanel: gamePanel}
}
