package model

import (
	"container/list"
	"image/color"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

var OrangeColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}

type Bullet struct {
	*Sprite
}

func NewBullet(falcon common.IFalcon) *Bullet {
	bullet := &Bullet{
		Sprite: NewSprite(),
	}

	bullet.team = common.FRIEND
	bullet.color = OrangeColor
	bullet.expiry = 20
	bullet.radius = 6

	bullet.center = falcon.GetCenter()
	bullet.orientation = falcon.GetOrientation()

	const (
		FIRE_POWER        = 35.0
		KICK_BACK_DIVISOR = 36.0
	)

	vectorX := math.Cos(bullet.orientation*math.Pi/180) * FIRE_POWER
	vectorY := math.Sin(bullet.orientation*math.Pi/180) * FIRE_POWER

	bullet.deltaX = falcon.GetDeltaX() + vectorX
	bullet.deltaY = falcon.GetDeltaY() + vectorY

	falcon.SetDeltaX(falcon.GetDeltaX() - vectorX/KICK_BACK_DIVISOR)
	falcon.SetDeltaY(falcon.GetDeltaY() - vectorY/KICK_BACK_DIVISOR)

	listPoints := []prime.Point{
		{X: 0, Y: 3},
		{X: 1, Y: -1},
		{X: 0, Y: 0},
		{X: -1, Y: -1},
	}

	bullet.cartesians = listPoints

	return bullet
}

func (b *Bullet) Move() {
	b.move(b)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.renderVector(screen)
}

func (b *Bullet) GetCenter() prime.Point {
	return b.center
}

func (b *Bullet) GetRadius() int {
	return b.radius
}

func (b *Bullet) GetTeam() common.Team {
	return b.team
}

func (b *Bullet) AddToGame(list *list.List) {
	b.addToGame(list, b)
}

func (b *Bullet) RemoveFromGame(list *list.List) {
	b.removeFromGame(list, b)

	common.PlaySound("thump.wav")
}