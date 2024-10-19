package model

import (
	"container/list"
	"image/color"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	center      prime.Point
	deltaX      float64
	deltaY      float64
	team        common.Team
	radius      int
	orientation float64
	expiry      int
	spin        float64
	cartesians  []prime.Point
	color       color.Color
	//private Map<?, BufferedImage> rasterMap;

	commandCenter common.ICommandCenter
}

func NewSprite(commandCenter common.ICommandCenter) *Sprite {
	return &Sprite{
		center: prime.Point{
			X: float64(common.GenerateRandomInt(common.DIM.Width)),
			Y: float64(common.GenerateRandomInt(common.DIM.Height)),
		},
		commandCenter: commandCenter,
	}
}

func (s *Sprite) move(movable common.Movable) {
	scalarX := s.commandCenter.GetUniDim().Width
	scalarY := s.commandCenter.GetUniDim().Height

	if s.center.X > float64(scalarX*common.DIM.Width) {
		s.center.X = 1
	} else if s.center.X < 0 {
		s.center.X = float64(scalarX*common.DIM.Width) - 1
	} else if s.center.Y > float64(scalarY*common.DIM.Height) {
		s.center.Y = 1
	} else if s.center.Y < 0 {
		s.center.Y = float64(scalarY*common.DIM.Height) - 1
	} else {
		newXPos := s.center.X
		newYPos := s.center.Y

		//if GetCommandCenterInstance().FalconPositionFixed {
		//	newXPos -= GetCommandCenterInstance().Falcon.DeltaX
		//	newYPos -= GetCommandCenterInstance().Falcon.DeltaY
		//}

		s.center.X = math.Round(newXPos + s.deltaX)
		s.center.Y = math.Round(newYPos + s.deltaX)
	}

	if s.expiry > 0 {
		s.expire(movable)
	}

	if s.spin != 0 {
		s.orientation += s.spin
	}
}

func (s *Sprite) expire(movable common.Movable) {
	if s.expiry == 1 {
		s.commandCenter.GetGameOpsQueue().Enqueue(movable, common.REMOVE)
		return
	}
	s.expiry--
}

func (s *Sprite) somePosNegValue(seed float64) float64 {
	randomNumber := common.GenerateRandomFloat64(seed)
	if common.GenerateRandomInt(2) == 0 {
		return randomNumber
	}
	return -randomNumber
}

func (s *Sprite) renderVector(screen *ebiten.Image) {
	polars := common.CartesiansToPolars(s.cartesians)

	var points []prime.Point
	for _, p := range polars {
		rotated := common.RotatePolarByOrientation(p, s.orientation)
		cartesian := common.PolarToCartesian(rotated, s.radius)
		adjusted := common.AdjustForLocation(cartesian, s.center)
		points = append(points, adjusted)
	}

	common.DrawPolygon(screen, points, s.color)
}

func (s *Sprite) addToGame(list *list.List, movable common.Movable) {
	list.PushBack(movable)
}

func (s *Sprite) removeFromGame(list *list.List, movable common.Movable) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == movable {
			list.Remove(e)
			break
		}
	}
}
