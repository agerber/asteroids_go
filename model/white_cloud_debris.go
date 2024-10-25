package model

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

var SLOW_MO = int(math.Round(3))

type WhiteCloudDebris struct {
	*Sprite

	index int
}

func NewWhiteCloudDebris(explodingSprite *Sprite) *WhiteCloudDebris {
	whiteCloudDebris := &WhiteCloudDebris{
		Sprite: explodingSprite,
	}

	whiteCloudDebris.team = common.DEBRIS

	whiteCloudDebris.rasterMap = make(map[interface{}]*ebiten.Image)
	whiteCloudDebris.rasterMap[0] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-1-column-1.png"))
	whiteCloudDebris.rasterMap[1] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-1-column-2.png"))
	whiteCloudDebris.rasterMap[2] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-1-column-3.png"))
	whiteCloudDebris.rasterMap[3] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-2-column-1.png"))
	whiteCloudDebris.rasterMap[4] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-2-column-2.png"))
	whiteCloudDebris.rasterMap[5] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-2-column-3.png"))
	whiteCloudDebris.rasterMap[6] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-3-column-1.png"))
	whiteCloudDebris.rasterMap[7] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-3-column-2.png"))
	whiteCloudDebris.rasterMap[8] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/exp/row-3-column-3.png"))

	//expire it out after it has done its animation. Multiply by SLOW_MO to slow down the animation
	whiteCloudDebris.expiry = len(whiteCloudDebris.rasterMap) * SLOW_MO

	whiteCloudDebris.radius = int(math.Round(float64(explodingSprite.radius) * 1.3))

	return whiteCloudDebris
}

func (w *WhiteCloudDebris) Move() {
	w.move(w)
}

func (w *WhiteCloudDebris) Draw(screen *ebiten.Image) {
	w.renderRaster(screen, w.rasterMap[w.index])
	if w.expiry%SLOW_MO == 0 {
		w.index++
	}
}

func (w *WhiteCloudDebris) GetCenter() prime.Point {
	return w.center
}

func (w *WhiteCloudDebris) GetRadius() int {
	return w.radius
}

func (w *WhiteCloudDebris) GetTeam() common.Team {
	return w.team
}

func (w *WhiteCloudDebris) AddToGame(list *list.List) {
	w.addToGame(list, w)
}

func (w *WhiteCloudDebris) RemoveFromGame(list *list.List) {
	w.removeFromGame(list, w)
}
