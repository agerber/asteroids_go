package commandcenter

import "github.com/agerber/asteroids_go/model"

type Action int

const (
	ADD Action = iota
	REMOVE
)

type GameOp struct {
	Movable model.Movable
	Action  Action
}
