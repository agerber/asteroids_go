package model

import (
	"container/list"
	"image/color"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	center prime.Point
	color  color.RGBA
}

func NewStar() common.Movable {
	bright := uint8(common.GenerateRandomInt(226))

	return &Star{
		center: prime.Point{
			X: float64(common.GenerateRandomInt(common.DIM.Width)),
			Y: float64(common.GenerateRandomInt(common.DIM.Height)),
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

func (s *Star) GetTeam() common.Team {
	return common.DEBRIS
}

func (s *Star) Move() {
	if !common.GetCommandCenterInstance().IsFalconPositionFixed() {
		return
	}

	if s.center.X > float64(common.DIM.Width) {
		s.center.X = 1
	} else if s.center.X < 0 {
		s.center.X = float64(common.DIM.Width - 1)
	} else if s.center.Y > float64(common.DIM.Height) {
		s.center.Y = 1
	} else if s.center.Y < 0 {
		s.center.Y = float64(common.DIM.Height - 1)
	} else {
		//move star in opposite direction of falcon.
		s.center.X -= common.GetCommandCenterInstance().GetFalcon().GetDeltaX()
		s.center.Y -= common.GetCommandCenterInstance().GetFalcon().GetDeltaY()
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
