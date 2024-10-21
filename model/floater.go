package model

import (
	"image/color"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
)

type Floater struct {
	*Sprite
}

func NewFloater() *Floater {
	floater := &Floater{
		Sprite: NewSprite(),
	}

	floater.team = common.FLOATER
	floater.expiry = 250
	floater.color = color.White
	floater.radius = 50
	floater.spin = floater.somePosNegValue(2)
	floater.deltaX = floater.somePosNegValue(5)
	floater.deltaY = floater.somePosNegValue(5)

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
