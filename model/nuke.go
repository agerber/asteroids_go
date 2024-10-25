package model

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

var NUKE_EXPIRY = int(math.Round(60))

type Nuke struct {
	*Sprite

	nukeState int
}

func NewNuke(falcon common.IFalcon) *Nuke {
	nuke := &Nuke{
		Sprite:    NewSprite(),
		nukeState: 0,
	}

	nuke.team = common.FRIEND
	nuke.color = colornames.Yellow
	nuke.expiry = NUKE_EXPIRY
	nuke.radius = 0

	nuke.center = falcon.GetCenter()

	const FIRE_POWER = 11.0

	vectorX := math.Cos(falcon.GetOrientation()*math.Pi/180) * FIRE_POWER
	vectorY := math.Sin(falcon.GetOrientation()*math.Pi/180) * FIRE_POWER

	nuke.deltaX = falcon.GetDeltaX() + vectorX
	nuke.deltaY = falcon.GetDeltaY() + vectorY

	return nuke
}

func (n *Nuke) Move() {
	n.move(n)

	if n.expiry%(NUKE_EXPIRY/6) == 0 {
		n.nukeState++
	}

	switch n.nukeState {
	case 0:
		n.radius = 17
	case 1, 2, 3:
		n.radius += 8
	default:
		n.radius -= 11
	}
}

func (n *Nuke) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, float32(n.center.X), float32(n.center.Y), float32(n.radius), 1, n.color, false)
}

func (n *Nuke) GetCenter() prime.Point {
	return n.center
}

func (n *Nuke) GetRadius() int {
	return n.radius
}

func (n *Nuke) GetTeam() common.Team {
	return n.team
}

func (n *Nuke) AddToGame(list *list.List) {
	if common.GetCommandCenterInstance().GetFalcon().GetNukeMeter() > 0 {
		n.addToGame(list, n)
		common.PlaySound("nuke.wav")
		common.GetCommandCenterInstance().GetFalcon().SetNukeMeter(0)
	}
}

func (n *Nuke) RemoveFromGame(list *list.List) {
	if n.expiry == 0 {
		n.removeFromGame(list, n)
	}
}
