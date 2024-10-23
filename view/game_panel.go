package view

import (
	"container/list"
	"fmt"
	"image/color"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	SHIP_RADIUS = 15
)

var (
	OrangeColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}
)

type GamePanel struct {
	dim                 common.Dimension
	pointShipsRemaining []prime.Point
}

func NewGamePanel(dim common.Dimension) *GamePanel {
	// Robert Alef's awesome falcon design
	pointShipsRemaining := make([]prime.Point, 0, 36)
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 0, Y: 9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: 6})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: 3})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -4, Y: 1})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 4, Y: 1})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -4, Y: 1})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -4, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -4, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -10, Y: -8})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -5, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -7, Y: -11})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -4, Y: -11})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -2, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -2, Y: -10})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: -10})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: -1, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: -10})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 2, Y: -10})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 2, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 4, Y: -11})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 7, Y: -11})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 5, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 10, Y: -8})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 4, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: -9})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 4, Y: -2})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 4, Y: 1})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: 3})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 1, Y: 6})
	pointShipsRemaining = append(pointShipsRemaining, prime.Point{X: 0, Y: 9})

	return &GamePanel{
		dim:                 dim,
		pointShipsRemaining: pointShipsRemaining,
	}
}

func (g *GamePanel) Draw(screen *ebiten.Image) {
	g.drawNumFrame(screen)
	if common.GetCommandCenterInstance().IsGameOver() {
		g.displayTextOnScreen(screen, []string{
			"GAME OVER",
			"use the arrow keys to turn and thrust",
			"use the space bar to fire",
			"'S' to Start",
			"'P' to Pause",
			"'Q' to Quit",
			"'M' to toggle music",
			"'A' to toggle radar",
		})
	} else if common.GetCommandCenterInstance().IsPaused() {
		g.displayTextOnScreen(screen, []string{"Game Paused"})
	} else {
		g.moveDrawMovables(screen,
			common.GetCommandCenterInstance().GetMovDebris(),
			common.GetCommandCenterInstance().GetMovFriends(),
			common.GetCommandCenterInstance().GetMovFoes(),
			common.GetCommandCenterInstance().GetMovFloaters())
		g.drawNumberShipsRemaining(screen)
	}
}

func (g *GamePanel) moveDrawMovables(screen *ebiten.Image, teams ...*list.List) {
	for _, team := range teams {
		for e := team.Front(); e != nil; e = e.Next() {
			movable := e.Value.(common.Movable)
			movable.Move()
			movable.Draw(screen)
		}
	}
}

func (g *GamePanel) drawNumberShipsRemaining(screen *ebiten.Image) {
	numFalcons := common.GetCommandCenterInstance().GetNumFalcons()
	for i := numFalcons; i > 1; i-- {
		g.drawOneShip(screen, i)
	}
}

func (g *GamePanel) drawOneShip(screen *ebiten.Image, offset int) {
	XPos := g.dim.Width - (27 * offset)
	YPos := g.dim.Height - 45

	polars := common.CartesiansToPolars(g.pointShipsRemaining)

	rotatePolarBy90 := func(pp prime.PolarPoint) prime.PolarPoint {
		return prime.PolarPoint{
			R:     pp.R,
			Theta: pp.Theta + math.Pi/2,
		}
	}
	polarToCartesian := func(pp prime.PolarPoint) prime.Point {
		return prime.Point{
			X: pp.R * SHIP_RADIUS * math.Sin(pp.Theta),
			Y: pp.R * SHIP_RADIUS * math.Cos(pp.Theta),
		}
	}
	adjustForLocation := func(point prime.Point) prime.Point {
		return prime.Point{
			X: point.X + float64(XPos),
			Y: point.Y + float64(YPos),
		}
	}

	points := make([]prime.Point, len(polars))
	for i, pp := range polars {
		pp = rotatePolarBy90(pp)
		point := polarToCartesian(pp)
		point = adjustForLocation(point)
		points[i] = point
	}

	common.DrawPolygon(screen, points, OrangeColor)
}

// TODO: Update to new font
var normalFont = basicfont.Face7x13

func (g *GamePanel) drawNumFrame(screen *ebiten.Image) {
	numFrameText := fmt.Sprintf("FRAME[GO]:%d", common.GetCommandCenterInstance().GetFrame())
	text.Draw(screen, numFrameText, normalFont, normalFont.Width, common.DIM.Height-(normalFont.Height), color.White)
}

func (g *GamePanel) displayTextOnScreen(screen *ebiten.Image, lines []string) {
	var spacer int
	for _, str := range lines {
		bounds, _ := font.BoundString(normalFont, str)
		width := (bounds.Max.X - bounds.Min.X).Ceil()
		x := (screen.Bounds().Dx() - width) / 2
		spacer += 40
		y := screen.Bounds().Dy()/4 + normalFont.Height + spacer
		text.Draw(screen, str, normalFont, x, y, color.White)
	}
}
