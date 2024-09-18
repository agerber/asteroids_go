package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuView struct{}

func (v *MenuView) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "GAME OVER", 320-len("GAME OVER")*3, 100)
	ebitenutil.DebugPrintAt(screen, "use the arrow keys to turn and thrust", 320-len("use the arrow keys to turn and thrust")*3, 120)
	ebitenutil.DebugPrintAt(screen, "use the space to fire", 320-len("use the space to fire")*3, 140)
	ebitenutil.DebugPrintAt(screen, "'S' to Start", 320-len("'S' to Start")*3, 160)
	ebitenutil.DebugPrintAt(screen, "'P' to Pause", 320-len("'S' to Pause")*3, 180)
	ebitenutil.DebugPrintAt(screen, "'Q' to Quit", 320-len("'S' to Quit")*3, 200)
	ebitenutil.DebugPrintAt(screen, "'M' to toggle music", 320-len("'M' to toggle music")*3, 220)
	ebitenutil.DebugPrintAt(screen, "'A' to toggle radar", 320-len("'M' to toggle radar")*3, 240)
}

type GameView struct{}

func (v *GameView) Draw(screen *ebiten.Image) {
	// Implement game rendering here
}

type SettingsView struct{}

func (v *SettingsView) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Settings", 320-len("Settings")*3, 100)
	ebitenutil.DebugPrintAt(screen, "'B' to Back to Menu", 320-len("'B' to Back to Menu")*3, 150)
}
