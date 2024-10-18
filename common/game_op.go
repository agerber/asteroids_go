package common

type Action int

const (
	ADD Action = iota
	REMOVE
)

type GameOp struct {
	Movable Movable
	Action  Action
}
