package model

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var (
	BULLET_EXPIRY            = int(math.Round(20 * common.GOLANG_FRAMES_SCALE_FACTOR))
	BULLET_FIRE_POWER        = 35.0 / common.GOLANG_FRAMES_SCALE_FACTOR
	BULLET_KICK_BACK_DIVISOR = 36.0 * common.GOLANG_FRAMES_SCALE_FACTOR
)

type Bullet struct {
	*Sprite
}

func NewBullet(falcon common.IFalcon) *Bullet {
	bullet := &Bullet{
		Sprite: NewSprite(),
	}

	bullet.team = common.FRIEND
	bullet.color = colornames.Orange
	bullet.expiry = BULLET_EXPIRY
	bullet.radius = 6

	bullet.center = falcon.GetCenter()
	bullet.orientation = falcon.GetOrientation()

	vectorX := math.Cos(bullet.orientation*math.Pi/180) * BULLET_FIRE_POWER
	vectorY := math.Sin(bullet.orientation*math.Pi/180) * BULLET_FIRE_POWER

	bullet.deltaX = falcon.GetDeltaX() + vectorX
	bullet.deltaY = falcon.GetDeltaY() + vectorY

	falcon.SetDeltaX(falcon.GetDeltaX() - vectorX/BULLET_KICK_BACK_DIVISOR)
	falcon.SetDeltaY(falcon.GetDeltaY() - vectorY/BULLET_KICK_BACK_DIVISOR)

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

	common.PlaySound("thump.wav")
}

func (b *Bullet) RemoveFromGame(list *list.List) {
	b.removeFromGame(list, b)
}
