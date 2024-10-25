package model

import (
	"container/list"
	"image/color"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

const MINI_MAP_PERCENT = 0.31

var PumpkinColor = color.RGBA{R: 200, G: 100, B: 50, A: 255}

type MiniMap struct {
	*Sprite

	aspectRatio prime.AspectRatio
}

func NewMiniMap() *MiniMap {
	miniMap := &MiniMap{
		Sprite: NewSprite(),
	}

	miniMap.team = common.DEBRIS
	miniMap.center = prime.Point{X: 0, Y: 0}

	return miniMap
}

func (m *MiniMap) Move() {
	m.move(m)
}

func (m *MiniMap) Draw(screen *ebiten.Image) {
	if !common.GetCommandCenterInstance().IsRadar() {
		return
	}

	universeDim := common.GetCommandCenterInstance().GetUniDim()

	m.aspectRatio = aspectAdjustedRatio(universeDim)

	miniWidth := float32(MINI_MAP_PERCENT * float64(common.DIM.Width) * m.aspectRatio.Width)
	miniHeight := float32(MINI_MAP_PERCENT * float64(common.DIM.Height) * m.aspectRatio.Height)

	vector.StrokeRect(screen, 0, 0, miniWidth, miniHeight, 1, colornames.Darkgray, false)

	miniViewPortWidth := miniWidth / float32(universeDim.Width)
	miniViewPortHeight := miniHeight / float32(universeDim.Height)
	vector.StrokeRect(screen, 0, 0, miniViewPortWidth, miniViewPortHeight, 1, colornames.Darkgray, false)

	for e := common.GetCommandCenterInstance().GetMovDebris().Front(); e != nil; e = e.Next() {
		mov := e.Value.(common.Movable)
		translatedPoint := m.translatePoint(mov.GetCenter(), universeDim)
		vector.StrokeCircle(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 1, 1, colornames.Darkgray, false)
	}

	for e := common.GetCommandCenterInstance().GetMovFoes().Front(); e != nil; e = e.Next() {
		asteroid, ok := e.Value.(*Asteroid)
		if !ok {
			continue
		}
		translatedPoint := m.translatePoint(asteroid.GetCenter(), universeDim)
		switch asteroid.GetSize() {
		case 0:
			vector.DrawFilledCircle(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 3, colornames.Lightgray, false)
		case 1:
			vector.StrokeCircle(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 3, 1, colornames.Lightgray, false)
		default:
			vector.StrokeCircle(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 2, 1, colornames.Lightgray, false)
		}
	}

	for e := common.GetCommandCenterInstance().GetMovFloaters().Front(); e != nil; e = e.Next() {
		mov := e.Value.(common.Movable)
		translatedPoint := m.translatePoint(mov.GetCenter(), universeDim)
		clr := colornames.Cyan
		if _, ok := mov.(*NukeFloater); ok {
			clr = colornames.Yellow
		}
		vector.DrawFilledRect(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 4, 4, clr, false)
	}

	for e := common.GetCommandCenterInstance().GetMovFriends().Front(); e != nil; e = e.Next() {
		mov := e.Value.(common.Movable)
		clr := PumpkinColor
		if _, ok := mov.(*Falcon); ok && common.GetCommandCenterInstance().GetFalcon().GetShield() > 0 {
			clr = colornames.Cyan
		} else if _, ok := mov.(*Nuke); ok {
			clr = colornames.Yellow
		}
		translatedPoint := m.translatePoint(mov.GetCenter(), universeDim)
		vector.DrawFilledCircle(screen, float32(translatedPoint.X), float32(translatedPoint.Y), 2, clr, false)
	}
}

func (m *MiniMap) GetCenter() prime.Point {
	return m.center
}

func (m *MiniMap) GetRadius() int {
	return m.radius
}

func (m *MiniMap) GetTeam() common.Team {
	return m.team
}

func (m *MiniMap) AddToGame(list *list.List) {
	m.addToGame(list, m)
}

func (m *MiniMap) RemoveFromGame(list *list.List) {
	m.removeFromGame(list, m)
}

func (m *MiniMap) translatePoint(point prime.Point, universeDim common.Dimension) prime.Point {
	return prime.Point{
		X: MINI_MAP_PERCENT * point.X / float64(universeDim.Width) * m.aspectRatio.Width,
		Y: MINI_MAP_PERCENT * point.Y / float64(universeDim.Height) * m.aspectRatio.Height,
	}
}

func aspectAdjustedRatio(universeDim common.Dimension) prime.AspectRatio {
	if universeDim.Width == universeDim.Height {
		return prime.AspectRatio{Width: 1, Height: 1}
	} else if universeDim.Width > universeDim.Height {
		wMultiple := float64(universeDim.Width) / float64(universeDim.Height)
		aspectRatio := prime.AspectRatio{Width: wMultiple, Height: 1}
		aspectRatio.Scale(0.5)
		return aspectRatio
	}
	hMultiple := float64(universeDim.Height) / float64(universeDim.Width)
	aspectRatio := prime.AspectRatio{Width: 1, Height: hMultiple}
	aspectRatio.Scale(0.5)
	return aspectRatio
}
