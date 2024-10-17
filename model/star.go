package model

import (
	"container/list"
	"image/color"

	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/agerber/asteroids_go/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	center prime.Point
	color  color.RGBA
}

func NewStar() *Star {
	bright := uint8(utils.GenerateRandomInt(226))

	return &Star{
		center: prime.Point{
			X: float64(utils.GenerateRandomInt(config.DIM.Width)),
			Y: float64(utils.GenerateRandomInt(config.DIM.Height)),
		},
		color: color.RGBA{
			R: bright,
			G: bright,
			B: bright,
			A: 255,
		},
	}
}

func (s *Star) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, float32(s.center.X), float32(s.center.Y), float32(s.GetRadius()), 1, s.color, false)
}

func (s *Star) GetCenter() prime.Point {
	return s.center
}

func (s *Star) GetRadius() int {
	return 1
}

func (s *Star) GetTeam() Team {
	return DEBRIS
}

func (s *Star) Move() {
	//if (!CommandCenter.getInstance().isFalconPositionFixed()) return;

	if s.center.X > float64(config.DIM.Width) {
		s.center.X = 1
	} else if s.center.X < 0 {
		s.center.X = float64(config.DIM.Width - 1)
	} else if s.center.Y > float64(config.DIM.Height) {
		s.center.Y = 1
	} else if s.center.Y < 0 {
		s.center.Y = float64(config.DIM.Height - 1)
	} else {
		s.center.X += 1
		s.center.Y += 1
		//center.x = (int) Math.round(center.x - CommandCenter.getInstance().getFalcon().getDeltaX());
		//center.y = (int) Math.round(center.y - CommandCenter.getInstance().getFalcon().getDeltaY());
	}
}

func (s *Star) AddToGame(list *list.List) {
	list.PushBack(s)
}

func (s *Star) RemoveFromGame(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == s {
			list.Remove(e)
			break
		}
	}
}
