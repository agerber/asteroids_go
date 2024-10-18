package utils

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func GenerateRandomInt(bound int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(bound)
}

func CartesiansToPolars(pointCartesians []prime.Point) []prime.PolarPoint {
	hypotenuseOfPoint := func(point prime.Point) float64 {
		return math.Sqrt(math.Pow(point.X, 2) + math.Pow(point.Y, 2))
	}

	var largestHyp float64
	for _, point := range pointCartesians {
		hyp := hypotenuseOfPoint(point)
		if hyp > largestHyp {
			largestHyp = hyp
		}
	}

	cartToPolarTransform := func(point prime.Point, largestHyp float64) prime.PolarPoint {
		return prime.PolarPoint{
			R:     hypotenuseOfPoint(point) / largestHyp,
			Theta: math.Atan2(point.Y, point.X),
		}
	}

	polars := make([]prime.PolarPoint, 0, len(pointCartesians))
	for _, point := range pointCartesians {
		polars = append(polars, cartToPolarTransform(point, largestHyp))
	}

	return polars
}

func DrawPolygon(screen *ebiten.Image, points []prime.Point, color color.Color) {
	// Draw lines between each pair of points
	for i := 0; i < len(points); i++ {
		start := points[i]
		end := points[(i+1)%len(points)] // Wrap around to the first point after the last point
		vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 1, color, false)
	}
}
