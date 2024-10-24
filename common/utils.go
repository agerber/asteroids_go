package common

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomInt(bound int) int {
	return random.Intn(bound)
}

func GenerateRandomFloat64(bound float64) float64 {
	return random.Float64() * bound
}

func DistanceBetween2Points(point1 prime.Point, point2 prime.Point) float64 {
	distance := math.Sqrt(math.Pow(point1.X-point2.X, 2) + math.Pow(point1.Y-point2.Y, 2))
	return math.Abs(distance)
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

func PolarToCartesian(pp prime.PolarPoint, radius int) prime.Point {
	return prime.Point{
		X: pp.R * float64(radius) * math.Sin(pp.Theta),
		Y: pp.R * float64(radius) * math.Cos(pp.Theta),
	}
}

func RotatePolarByOrientation(pp prime.PolarPoint, orientation float64) prime.PolarPoint {
	return prime.PolarPoint{
		R:     pp.R,
		Theta: pp.Theta + orientation, // rotated Theta
	}
}

func AdjustForLocation(p prime.Point, center prime.Point) prime.Point {
	return prime.Point{
		X: center.X + p.X,
		Y: center.Y - p.Y,
	}
}

func DrawPolygon(screen *ebiten.Image, points []prime.Point, color color.Color) {
	// Draw lines between each pair of points
	for i := 0; i < len(points); i++ {
		start := points[i]
		end := points[(i+1)%len(points)] // Wrap around to the first point after the last point
		vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 1, color, false)
	}
}
