package controller

import (
	"image/color"
	"image/draw"
	"math"

	"github.com/agerber/asteroids_go/src/mvc/model/prime"
)

type Utils struct{}

// CartesiansToPolar converts a list of Cartesian points to a list of Polar points.
func (u *Utils) CartesiansToPolar(pntCartesians []prime.Point) []prime.PolarPoint {
	largestHypotenuse := 0.0
	for _, pnt := range pntCartesians {
		hypotenuse := math.Sqrt(math.Pow(float64(pnt.X), 2) + math.Pow(float64(pnt.Y), 2))
		if hypotenuse > largestHypotenuse {
			largestHypotenuse = hypotenuse
		}
	}

	polarPoints := make([]prime.PolarPoint, len(pntCartesians))
	for i, pnt := range pntCartesians {
		hypotenuse := math.Sqrt(math.Pow(float64(pnt.X), 2) + math.Pow(float64(pnt.Y), 2))
		polarPoints[i] = prime.PolarPoint{
			R:     hypotenuse / largestHypotenuse,
			Theta: radiansToDegrees(math.Atan2(float64(pnt.Y), float64(pnt.X))),
		}
	}
	return polarPoints
}

// Transparent makes an image transparent by adding an alpha channel if needed.
// ESSA FUNÇÃO VAI SER USADA NA CLASSE SPRITE
func (u *Utils) Transparent(img image.Image) image.Image {
	if img.ColorModel() != color.RGBAModel {
		img = draw.Drawer.Convert(img, "RGBA")
	}

	transparentImg := image.NewRGBA(img.Bounds())
	draw := image.NewDraw(transparentImg)
	draw.Draw(transparentImg, transparentImg.Bounds(), img, image.ZP, draw.Src)
	return transparentImg
}

// radiansToDegrees converts radians to degrees.
func radiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
