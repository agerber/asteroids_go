package prime

import "math"

type Point struct {
	X int
	Y int
}

// NewPoint creates a new Point instance.
func NewPoint(x int, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-other.X), 2) + math.Pow(float64(p.Y-other.Y), 2))

}

func (p Point) Clone() Point {
	return Point{X: p.X, Y: p.Y}
}
