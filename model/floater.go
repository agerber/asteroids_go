package model

import (
	"image/color"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
)

var FLOATER_EXPIRY = int(math.Round(250))

type Floater struct {
	*Sprite
}

func NewFloater() *Floater {
	floater := &Floater{
		Sprite: NewSprite(),
	}

	floater.team = common.FLOATER
	floater.expiry = FLOATER_EXPIRY
	floater.color = color.White
	floater.radius = 50
	floater.spin = floater.somePosNegValue(10)
	floater.deltaX = floater.somePosNegValue(10)
	floater.deltaY = floater.somePosNegValue(10)

	listPoints := []prime.Point{
		{X: 5, Y: 5},
		{X: 4, Y: 0},
		{X: 5, Y: -5},
		{X: 0, Y: -4},
		{X: -5, Y: -5},
		{X: -4, Y: 0},
		{X: -5, Y: 5},
		{X: 0, Y: 4},
	}

	floater.cartesians = listPoints

	return floater
}
