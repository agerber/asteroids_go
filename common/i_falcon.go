package common

import (
	"container/list"

	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

type IFalcon interface {
	Move()
	Draw(screen *ebiten.Image)
	GetCenter() prime.Point
	GetRadius() int
	GetTeam() Team
	AddToGame(list *list.List)
	RemoveFromGame(list *list.List)
	GetDeltaX() float64
	GetDeltaY() float64
}
