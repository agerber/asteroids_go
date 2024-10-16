package view

import (
	"image/color"
	"math"

	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/agerber/asteroids_go/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SHIP_RADIUS = 15
)

var (
	OrangeColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}
)

type GamePanel struct {
	dim                 config.Dimension
	pointShipsRemaining []prime.Point
}

func NewGamePanel(dim config.Dimension) *GamePanel {
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
	g.drawNumberShipsRemaining(screen)
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

	polars := utils.CartesiansToPolars(g.pointShipsRemaining)

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

	// Draw lines between each pair of points
	for i := 0; i < len(points); i++ {
		start := points[i]
		end := points[(i+1)%len(points)] // Wrap around to the first point after the last point
		vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 1, OrangeColor, false)
	}
}
