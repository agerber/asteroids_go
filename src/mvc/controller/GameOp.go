package controller

import "github.com/agerber/asteroids_go/src/mvc/model"

// GameOp represents a game operation (add or remove).
type GameOp struct {
	Action  GameOpAction
	Movable model.Movable
}

// GameOpAction represents the type of game operation.
type GameOpAction int

const (
	GameOpAdd    GameOpAction = 1
	GameOpRemove GameOpAction = 2
)
