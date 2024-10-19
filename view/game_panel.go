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

	commandCenter common.ICommandCenter
}

func NewGamePanel(dim common.Dimension, commandCenter common.ICommandCenter) *GamePanel {
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
		commandCenter:       commandCenter,
	}
}

func (g *GamePanel) Draw(screen *ebiten.Image) {
	g.drawNumFrame(screen)
	g.moveDrawMovables(screen,
		g.commandCenter.GetMovDebris(),
		g.commandCenter.GetMovFriends(),
		g.commandCenter.GetMovFoes(),
		g.commandCenter.GetMovFloaters())
	g.drawNumberShipsRemaining(screen)
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
	//int numFalcons = CommandCenter.getInstance().getNumFalcons(); TODO: convert it
	numFalcons := 36
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
	numFrameText := fmt.Sprintf("FRAME[GO]:%d", g.commandCenter.GetFrame())
	text.Draw(screen, numFrameText, normalFont, normalFont.Width, common.DIM.Height-(normalFont.Height), color.White)
}
