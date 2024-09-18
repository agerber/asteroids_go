package view

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	"github.com/agerber/asteroids_go/cmd/model"
)

// GamePanel represents the main game view.
type GamePanel struct {
	fontSmall font.Face
	fontBig   font.Face
	falconPts []model.PolarPoint
}

// NewGamePanel initializes a new GamePanel.
func NewGamePanel() *GamePanel {
	// Create points for falcon ship
	falconPts := []model.PolarPoint{
		{R: 10, Theta: 0},
		{R: 8, Theta: math.Pi / 4},
		{R: 10, Theta: math.Pi / 2},
		// Add all points from your design...
	}
	return &GamePanel{
		fontSmall: basicfont.Face7x13,
		fontBig:   basicfont.Face7x13, // Use a larger font for larger text
		falconPts: falconPts,
	}
}

// Update handles the game logic, called every frame.
func (gp *GamePanel) Update() error {
	cc := model.GetCommandCenter()
	if cc.IsGameOver {
		return nil
	}

	// Handle game logic here, like moving objects
	// controller.CommandCenter().MoveObjects()
	return nil
}

// Draw draws the game screen.
func (gp *GamePanel) Draw(screen *ebiten.Image) {
	// Fill the screen with black
	screen.Fill(color.Black)

	cc := model.GetCommandCenter()
	if cc.IsGameOver {
		gp.displayTextOnScreen(screen, "GAME OVER", "Press S to Start", "Press Q to Quit")
		return
	}

	if cc.IsPaused {
		gp.displayTextOnScreen(screen, "Game Paused")
		return
	}

	// Draw Falcon status, score, level, etc.
	gp.drawFalconStatus(screen)
	gp.drawMeters(screen)
	gp.drawShipsRemaining(screen)
}

// drawFalconStatus draws the Falcon status (level, score, etc.) in the top right.
func (gp *GamePanel) drawFalconStatus(screen *ebiten.Image) {
	cc := model.GetCommandCenter()

	levelText := fmt.Sprintf("Level: [%d]", cc.Level)
	scoreText := fmt.Sprintf("Score: %d", cc.Score)

	text.Draw(screen, levelText, gp.fontSmall, 400, 20, color.White)
	text.Draw(screen, scoreText, gp.fontSmall, 400, 40, color.White)

	// Status messages like "SLOW DOWN" or "PRESS F for NUKE"
	statusMessages := []string{}
	if cc.Falcon.IsMaxSpeedAttained {
		statusMessages = append(statusMessages, "WARNING - SLOW DOWN")
	}
	if cc.Falcon.NukeMeter > 0 {
		statusMessages = append(statusMessages, "PRESS F for NUKE")
	}

	if len(statusMessages) > 0 {
		gp.displayTextOnScreen(screen, statusMessages...)
	}
}

// drawMeters draws the shield and nuke meters.
func (gp *GamePanel) drawMeters(screen *ebiten.Image) {
	cc := model.GetCommandCenter()

	shieldValue := cc.Falcon.Shield / 2  // Scale to 0-100
	nukeValue := cc.Falcon.NukeMeter / 6 // Scale to 0-100

	gp.drawOneMeter(screen, color.RGBA{0, 255, 255, 255}, 1, shieldValue)
	gp.drawOneMeter(screen, color.RGBA{255, 255, 0, 255}, 2, nukeValue)
}

// drawOneMeter draws a single meter (e.g., shield or nuke).
func (gp *GamePanel) drawOneMeter(screen *ebiten.Image, col color.Color, offset int, percent int) {
	xVal := 400 + (100 * offset)
	yVal := 450

	// Background of meter
	rect := ebiten.NewImage(100, 10)
	rect.Fill(color.RGBA{50, 50, 50, 255}) // Dark Gray

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xVal), float64(yVal))
	screen.DrawImage(rect, op)

	// Draw the actual meter fill
	rect.Fill(col)
	screen.DrawImage(rect.SubImage(image.Rect(0, 0, percent, 10)).(*ebiten.Image), op)
}

// drawShipsRemaining draws the remaining Falcon icons.
func (gp *GamePanel) drawShipsRemaining(screen *ebiten.Image) {
	cc := model.GetCommandCenter()
	for i := 0; i < cc.NumFalcons; i++ {
		gp.drawOneShip(screen, i)
	}
}

// drawOneShip draws a single Falcon icon at a given offset.
func (gp *GamePanel) drawOneShip(screen *ebiten.Image, offset int) {
	xPos := 600 - offset*30
	yPos := 400

	// Convert the Falcon points from polar to Cartesian and draw
	vertices := []ebiten.Vertex{}
	for _, polar := range gp.falconPts {
		x, y := PolarToCartesian(polar)
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(float32(xPos) + x),
			DstY:   float32(float32(yPos) + y),
			ColorR: 1, ColorG: 1, ColorB: 0, ColorA: 1,
		})
	}

	indices := []uint16{0, 1, 2} // Triangle indices for the ship shape
	screen.DrawTriangles(vertices, indices, ebiten.NewImage(1, 1), nil)
}

// displayTextOnScreen draws text centered on the screen.
func (gp *GamePanel) displayTextOnScreen(screen *ebiten.Image, lines ...string) {
	for i, line := range lines {
		bounds := text.BoundString(gp.fontSmall, line)
		x := (640 - bounds.Dx()) / 2 // Center the text horizontally
		y := (480 / 4) + i*40        // Stack the lines with some spacing
		text.Draw(screen, line, gp.fontSmall, x, y, color.White)
	}
}

// Layout defines the window size for Ebiten.
func (gp *GamePanel) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func PolarToCartesian(polar model.PolarPoint) (x, y float32) {
	x = float32(polar.R * math.Cos(polar.Theta))
	y = float32(polar.R * math.Sin(polar.Theta))
	return x, y
}
