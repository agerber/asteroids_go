package model

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const SPAWN_SHIELD_FLOATER = common.FRAMES_PER_SECOND * 25

var SHIELD_FLOATER_EXPIRY = int(math.Round(260))

type ShieldFloater struct {
	*Floater
}

func NewShieldFloater() common.Movable {
	shieldFloater := &ShieldFloater{
		Floater: NewFloater(),
	}

	shieldFloater.color = colornames.Cyan
	shieldFloater.expiry = SHIELD_FLOATER_EXPIRY

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
	removed := s.Floater.removeFromGame(list, s)
	if !removed {
		return
	}

	if s.expiry > 0 {
		common.PlaySound("shieldup.wav")
		common.GetCommandCenterInstance().GetFalcon().SetShield(FALCON_MAX_SHIELD)
	}
}
