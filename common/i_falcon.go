package common

import (
	"container/list"

	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

type TurnState int

const (
	IDLE TurnState = iota
	LEFT
	RIGHT
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
	DecrementFalconNumAndSpawn()
	SetThrusting(thrusting bool)
	SetTurnState(turnState TurnState)
	GetOrientation() float64
	SetDeltaX(deltaX float64)
	SetDeltaY(deltaY float64)
	GetNukeMeter() int
	SetNukeMeter(nukeMeter int)
}
