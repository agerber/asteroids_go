package utils

import (
	"math"
	"math/rand"
	"time"

	"github.com/agerber/asteroids_go/model/prime"
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
