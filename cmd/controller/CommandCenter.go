package controller

import (
	"github.com/agerber/asteroids_go/cmd/model"
	"github.com/agerber/asteroids_go/cmd/view"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameController manages the game's state and updates.
type GameController struct {
	gamePanel *view.GamePanel
}

// NewGameController initializes a new game controller.
func NewGameController() *GameController {
	return &GameController{
		gamePanel: view.NewGamePanel(),
	}
}

// Update handles the game state transitions and updates the game logic.
func (gc *GameController) Update() error {
	cc := model.GetCommandCenter()

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyS):
		cc.IsGameOver = false
		cc.IsPaused = false
		cc.NumFalcons = 3
		cc.Score = 0
		cc.Level = 1
		cc.Falcon = &model.Falcon{Shield: 100, NukeMeter: 0}
	case ebiten.IsKeyPressed(ebiten.KeyP):
		cc.IsPaused = !cc.IsPaused
	case ebiten.IsKeyPressed(ebiten.KeyQ):
		// Add code to quit the game if needed
	}

	// Only update the game if it's not paused or over
	if !cc.IsPaused && !cc.IsGameOver {
		// Update game objects and logic here
	}

	// End the game if no falcons are left
	if cc.NumFalcons <= 0 {
		cc.IsGameOver = true
	}

	return gc.gamePanel.Update()
}

func (gc *GameController) Draw(screen *ebiten.Image) {
	gc.gamePanel.Draw(screen)
}

func (gc *GameController) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gc.gamePanel.Layout(outsideWidth, outsideHeight)
}
