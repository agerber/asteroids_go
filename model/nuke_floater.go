package model

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const SPAWN_NUKE_FLOATER = common.FRAMES_PER_SECOND * 12

var NUKE_FLOATER_EXPIRY = int(math.Round(120))

type NukeFloater struct {
	*Floater
}

func NewNukeFloater() common.Movable {
	nukeFloater := &NukeFloater{
		Floater: NewFloater(),
	}

	nukeFloater.color = colornames.Yellow
	nukeFloater.expiry = NUKE_FLOATER_EXPIRY

	return nukeFloater
}

func (n *NukeFloater) Move() {
	n.Floater.move(n)
}

func (n *NukeFloater) Draw(screen *ebiten.Image) {
	n.Floater.renderVector(screen)
}

func (n *NukeFloater) GetCenter() prime.Point {
	return n.Floater.center
}

func (n *NukeFloater) GetRadius() int {
	return n.Floater.radius
}

func (n *NukeFloater) GetTeam() common.Team {
	return n.Floater.team
}

func (n *NukeFloater) AddToGame(list *list.List) {
	n.Floater.addToGame(list, n)
}

func (n *NukeFloater) RemoveFromGame(list *list.List) {
	removed := n.Floater.removeFromGame(list, n)
	if !removed {
		return
	}

	if n.expiry > 0 {
		common.PlaySound("nuke-up.wav")
		common.GetCommandCenterInstance().GetFalcon().SetNukeMeter(FALCON_MAX_NUKE)
	}
}
