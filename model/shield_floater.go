package model

import (
	"container/list"
	"image/color"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

var CyanColor = color.RGBA{R: 0, G: 255, B: 255, A: 255}

type ShieldFloater struct {
	*Floater
}

func NewShieldFloater() common.Movable {
	shieldFloater := &ShieldFloater{
		Floater: NewFloater(),
	}

	shieldFloater.color = CyanColor
	shieldFloater.expiry = 600

	return shieldFloater
}

func (s *ShieldFloater) Move() {
	s.Floater.move(s)
}

func (s *ShieldFloater) Draw(screen *ebiten.Image) {
	s.Floater.renderVector(screen)
}

func (s *ShieldFloater) GetCenter() prime.Point {
	return s.Floater.center
}

func (s *ShieldFloater) GetRadius() int {
	return s.Floater.radius
}

func (s *ShieldFloater) GetTeam() common.Team {
	return s.Floater.team
}

func (s *ShieldFloater) AddToGame(list *list.List) {
	s.Floater.addToGame(list, s)
}

func (s *ShieldFloater) RemoveFromGame(list *list.List) {
	s.Floater.removeFromGame(list, s)

	if s.expiry > 0 {
		common.PlaySound("shieldup.wav")
		//CommandCenter.getInstance().getFalcon().setShield(Falcon.MAX_SHIELD)
	}
}
