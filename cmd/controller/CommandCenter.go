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

	// Handle input for Falcon
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		cc.Falcon.TurnState = model.LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		cc.Falcon.TurnState = model.RIGHT
	} else {
		cc.Falcon.TurnState = model.IDLE
	}

	cc.Falcon.Thrusting = ebiten.IsKeyPressed(ebiten.KeyArrowUp)

	// Start game with 'S'
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cc.IsGameOver = false
		cc.IsPaused = false
		cc.NumFalcons = 3
		cc.Score = 0
		cc.Level = 1
		cc.Falcon = model.NewFalcon()
	}

	// Pause game with 'P'
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		cc.IsPaused = !cc.IsPaused
	}

	// End game when falcons run out
	if cc.NumFalcons <= 0 {
		cc.IsGameOver = true
	}

	// Update game objects if not paused or over
	if !cc.IsPaused && !cc.IsGameOver {
		cc.Falcon.Move() // Update Falcon position and state
	}

	return gc.gamePanel.Update()
}

// Draw renders the game objects to the screen
func (gc *GameController) Draw(screen *ebiten.Image) {
	cc := model.GetCommandCenter()

	// Draw Falcon
	cc.Falcon.Draw(screen)

	// Draw other game objects from the game panel
	gc.gamePanel.Draw(screen)
}

func (gc *GameController) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gc.gamePanel.Layout(outsideWidth, outsideHeight)
}
